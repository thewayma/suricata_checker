package g
import (
    "io"
    "fmt"
    "sort"
    "strings"
    "crypto/md5"
)

// code == 0 => success
// code == 1 => bad request
type SimpleRpcResponse struct {
    Code int `json:"code"`
}

type Strategy struct {
    Id         int               `json:"id"`
    Metric     string            `json:"metric"`
    Tags       map[string]string `json:"tags"`
    Func       string            `json:"func"`       // e.g. max(#3) all(#3)
    Operator   string            `json:"operator"`   // e.g. < !=
    RightValue float64           `json:"rightValue"` // critical value
    MaxStep    int               `json:"maxStep"`
    Priority   int               `json:"priority"`
    Note       string            `json:"note"`
    Tpl        *Template         `json:"tpl"`
}

func (this *Strategy) String() string {
    return fmt.Sprintf(
        "<Id:%d, Metric:%s, Tags:%v, %s%s%s MaxStep:%d, P%d, %s, %v>",
        this.Id,
        this.Metric,
        this.Tags,
        this.Func,
        this.Operator,
        this.RightValue,
        this.MaxStep,
        this.Priority,
        this.Note,
        this.Tpl,
    )
}

type Template struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    ParentId int    `json:"parentId"`
    ActionId int    `json:"actionId"`
    Creator  string `json:"creator"`
}

func (this *Template) String() string {
    return fmt.Sprintf(
        "<Id:%d, Name:%s, ParentId:%d, ActionId:%d, Creator:%s>",
        this.Id,
        this.Name,
        this.ParentId,
        this.ActionId,
        this.Creator,
    )
}

type Event struct {
    Id          string            `json:"id"`
    Strategy    *Strategy         `json:"strategy"`
    //Expression  *Expression       `json:"expression"` //!< 废除expression机制
    Status      string            `json:"status"`
    Endpoint    string            `json:"endpoint"`
    LeftValue   float64           `json:"leftValue"`
    CurrentStep int               `json:"currentStep"`
    EventTime   int64             `json:"eventTime"`
    PushedTags  map[string]string `json:"pushedTags"`
}

func (this *Event) String() string {
    return fmt.Sprintf(
        "<Endpoint:%s, Status:%s, Strategy:%v, LeftValue:%s, CurrentStep:%d, PushedTags:%v>",
        this.Endpoint,
        this.Status,
        this.Strategy,
        //this.Expression,
        this.LeftValue,
        this.CurrentStep,
        this.PushedTags,
    )
}

func (this *Event) Priority() int {
    if this.Strategy != nil {
        return this.Strategy.Priority
    }
    return 1
}

type JudgeItem struct {
    Endpoint  string            `json:"endpoint"`
    Metric    string            `json:"metric"`
    Value     float64           `json:"value"`
    Timestamp int64             `json:"timestamp"`
    JudgeType string            `json:"judgeType"`
    Tags      map[string]string `json:"tags"`
}

func (this *JudgeItem) String() string {
    return fmt.Sprintf("<Endpoint:%s, Metric:%s, Value:%f, Timestamp:%d, JudgeType:%s Tags:%v>",
        this.Endpoint,
        this.Metric,
        this.Value,
        this.Timestamp,
        this.JudgeType,
        this.Tags)
}

func sortedTags(tags map[string]string) string {
    if tags == nil {
        return ""
    }

    size := len(tags)

    if size == 0 {
        return ""
    }

    if size == 1 {
        for k, v := range tags {
            return fmt.Sprintf("%s=%s", k, v)
        }
    }

    keys := make([]string, size)
    i := 0
    for k := range tags {
        keys[i] = k
        i++
    }

    sort.Strings(keys)

    ret := make([]string, size)
    for j, key := range keys {
        ret[j] = fmt.Sprintf("%s=%s", key, tags[key])
    }

    return strings.Join(ret, ",")
}

func Md5(raw string) string {
    h := md5.New()
    io.WriteString(h, raw)

    return fmt.Sprintf("%x", h.Sum(nil))
}

func (r *JudgeItem) PK() string {
    tags     := r.Tags
    endpoint := r.Endpoint
    metric   := r.Metric

    if tags == nil || len(tags) == 0 {
        return fmt.Sprintf("%s/%s", endpoint, metric)
    }
    return fmt.Sprintf("%s/%s/%s", endpoint, metric, sortedTags(tags))
}

func (this *JudgeItem) PrimaryKey() string {
    return Md5(this.PK())
}

type HistoryData struct {
    Timestamp int64   `json:"timestamp"`
    Value     float64 `json:"value"`
}
