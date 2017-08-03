package g
import (
   "fmt"
)

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

