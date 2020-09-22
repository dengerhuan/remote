package authz

var db = make(map[string]Vehicle)

type Vehicle struct {
	seq string
	vid string
}

func (v *Vehicle) GetId() string {
	return v.vid
}

func init() {
	db["0c23a8bf29a191f18aee814737e2a6ec"] = Vehicle{
		seq: "0c23a8bf29a191f18aee814737e2a6ec",
		vid: "VIN30000S20003",
	}

	db["dfd5503d9924a0d840d6ff6cc329c4cb"] = Vehicle{seq: "dfd5503d9924a0d840d6ff6cc329c4cb", vid: "001"}

	db["c36f08c814202504345b580264f4a874"]=Vehicle{seq:"c36f08c814202504345b580264f4a874",vid:"0001000001"}
}

func GetVehicleBySeq(seq string) (vid Vehicle, ob bool) {
	vid, ok := db[seq]
	return vid, ok
}
