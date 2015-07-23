package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmervine/sns2hipchat/sns"
	"strings"
	"text/template"
)

const CLOUDWATCH_ALARM_TEMPLATE = `<h4>New Alarm State: "{{.NewStateValue}}"</h4>
<hr />
<ul>
    <li><b>Name:</b> {{.AlarmName}}</li>
    <li><b>Reason:</b> {{.NewStateReason}}</li>
    <li><b>Previous:</b> {{.OldStateValue}}</li>
    <li><b>Metric:</b> {{.Trigger.MetricName}}</li>
    <li><b>Dimensions:</b>
        <ul>{{ range $k, $v := .Trigger.Dimensions }}
            <li><b>{{.Name}}:</b> {{.Value}}</li>{{end}}
        </ul>
    </li>
    <li><b>AWS Acct:</b> {{.AWSAccountID}}</li>
</ul>
<hr />
`

// Payload for JSON parsing
//
// - Commenting out that which can be ignored
// - Generated with: http://json2struct.mervine.net/
type cloudWatchAlarmPayload struct {
	AWSAccountID string `json:"AWSAccountId"`
	//AlarmDescription interface{} `json:"AlarmDescription"`
	AlarmName      string `json:"AlarmName"`
	NewStateReason string `json:"NewStateReason"`
	NewStateValue  string `json:"NewStateValue"`
	OldStateValue  string `json:"OldStateValue"`
	//Region          string `json:"Region"`
	//StateChangeTime string `json:"StateChangeTime"`
	Trigger struct {
		//ComparisonOperator string `json:"ComparisonOperator"`
		Dimensions []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"Dimensions"`
		//EvaluationPeriods int         `json:"EvaluationPeriods"`
		MetricName string `json:"MetricName"`
		//Namespace         string      `json:"Namespace"`
		//Period            int         `json:"Period"`
		//Statistic         string      `json:"Statistic"`
		//Threshold         int         `json:"Threshold"`
		//Unit              interface{} `json:"Unit"`
	} `json:"Trigger"`
}

func (p cloudWatchAlarmPayload) dimensions() string {
	var out []string
	for _, dim := range p.Trigger.Dimensions {
		out = append(out, fmt.Sprintf("- %s: %30s", dim.Name, dim.Value))
	}

	return strings.Join(out, "\n")
}

type CloudWatchAlarm struct {
	Formatter
}

func (f CloudWatchAlarm) decode(n *sns.Notification) (msg string, err error) {
	var p cloudWatchAlarmPayload

	dec := json.NewDecoder(strings.NewReader(n.Message))
	err = dec.Decode(&p)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	t, err := template.New("fmt").Parse(CLOUDWATCH_ALARM_TEMPLATE)
	if err != nil {
		panic(err)
	}

	err = t.Execute(&doc, p)
	if err != nil {
		panic(err)
	}

	msg = doc.String()
	return
}

func (f CloudWatchAlarm) Format(n *sns.Notification) (msg string, err error) {
	if p := recover(); p != nil {
		fmt.Printf("at=formatter type=alarm error=%s took=%v\n",
			p.(error).Error())
		fmtr := new(Raw)
		msg, err = fmtr.Format(n)
	}
	return f.decode(n)
}

func (f CloudWatchAlarm) FormatHTML(n *sns.Notification) (msg string, err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("at=formatter type=alarm error=%s took=%v\n",
				p.(error).Error())
			fmtr := new(Raw)
			msg, err = fmtr.FormatHTML(n)
		}
	}()
	return f.decode(n)
}
