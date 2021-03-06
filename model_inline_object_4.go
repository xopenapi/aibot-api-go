/*
 * aibot open api
 *
 * aibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// 操作类型，必填 PAUSE:暂停 START:开始
	Operation string `json:"operation"`
	// 任务id集合
	RobotCallJobIdSet []string `json:"robotCallJobIdSet,omitempty"`
	// 任务状态，NOT_STARTED(0, \"未开始\"),IN_PROCESS(1, \"进行中\")，COMPLETED(2, \"已完成\"),RUNNABLE(3, \"可运行\"),USER_PAUSE(4, \"用户暂停\"),SYSTEM_SUSPENDED(5, \"系统暂停\")，TERMINATE(6, \"已终止\"),IN_QUEUE(7, \"排队中\"),SYSTEM_HANG_UP(10, \"系统挂起\"),WAITING_FOR_REDIAL(11, \"等待重呼\"),ACCOUNT_DISABLE(12, \"账户禁用\"),MAINTAIN(13, \"系统维护\");
	StatusSet string `json:"statusSet,omitempty"`
	// 任务名称
	Name string `json:"name,omitempty"`
	// 开始任务创建日期
	BeginCreateDate string `json:"beginCreateDate,omitempty"`
	// 结束任务创建日期
	EndCreateDate string `json:"endCreateDate,omitempty"`
}
