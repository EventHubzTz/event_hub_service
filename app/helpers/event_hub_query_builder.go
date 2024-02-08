package helpers

var EventHubQueryBuilder = newEventHubQueryBuilder()

type eventHubQueryBuilder struct {
}

func newEventHubQueryBuilder() eventHubQueryBuilder {
	return eventHubQueryBuilder{}
}

func (_ eventHubQueryBuilder) QueryMicroServiceRequestIDActiveKey() string {
	return "SELECT t1.id,t1.request_id FROM event_hub_request_ids t1 " +
		"ORDER BY t1.id DESC LIMIT 1"
}

func (_ eventHubQueryBuilder) QueryGetUsers(role string) string {
	return "SELECT t1.id as id,t1.first_name,t1.last_name," +
		"t1.phone_number,t1.gender,t1.role,t1.active," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at," +
		"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at " +
		"FROM event_hub_users t1 " +
		"WHERE t1.active = true and t1.role = ? " +
		"ORDER BY t1.id DESC "
}

func (_ eventHubQueryBuilder) QueryUserDetails() string {
	return "SELECT t1.id as id,t1.first_name,t1.last_name," +
		"t1.phone_number,t1.gender,t1.role,t1.active," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at " +
		"FROM event_hub_users t1 " +
		"WHERE t1.active = true and t1.id = ?"
}
