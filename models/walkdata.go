package models

import (
	"fmt"
	// "errors"
	// "fmt"
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	// "reflect"
	// "strings"
	// "common/timeutils"
)

type Walkdata struct {
	Aliuid    string  `orm:"size(128)" json:"aliuid"`
	Step      float64 `json:"step"`
	Energy    float64 `json:"energy"`
	Distance  float64 `json:"distance"`
	Duration  float64 `json:"duration"`
	Timestamp int64   `json:"timestamp"`
	Datestamp int64   `json:"datestamp"`
}

var (
	tsclient          *tablestore.TableStoreClient
	walkDataTableName string
)

const (
	PKALIUID     = "pkaliuid"
	PKDATESTAMP  = "pkdatestamp"
	COLSTEP      = "step"
	COLENERGY    = "energy"
	COLDISTANCE  = "distance"
	COLDURATION  = "duration"
	COLTIMESTAMP = "timestamp"
	RETRY        = 3
)

func init() {
	endpoint := os.Getenv("TS_ALISPORTDATACENTER_ENDPOINT")
	instanceName := os.Getenv("TS_ALISPORTDATACENTER_INSTANCENAME")
	accessKeyId := os.Getenv("TS_ALISPORTDATACENTER_KEYID")
	accessKeySecret := os.Getenv("TS_ALISPORTDATACENTER_SECRET")
	walkDataTableName = os.Getenv("TS_WALKDATA_TABLENAME")
	tsclient = tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)
}

func addWalkdataWithGeneralTS(m *Walkdata) (err error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn(PKALIUID, m.Aliuid)
	putPk.AddPrimaryKeyColumn(PKDATESTAMP, m.Datestamp)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = walkDataTableName
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1

	getResp, _ := tsclient.GetRow(getRowRequest)
	colmap := getResp.GetColumnMap()
	if colmap != nil {
		columstep, ok := colmap.Columns[COLSTEP]
		if ok == true {
			step, _ := columstep[0].Value.(float64)
			if step >= m.Step {
				return nil
			}
		}
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = walkDataTableName
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn(COLSTEP, m.Step)
	putRowChange.AddColumn(COLENERGY, m.Energy)
	putRowChange.AddColumn(COLDISTANCE, m.Distance)
	putRowChange.AddColumn(COLDURATION, m.Duration)
	putRowChange.AddColumn(COLTIMESTAMP, m.Timestamp)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	// todo:添加重试逻辑
	_, err = tsclient.PutRow(putRowRequest)
	return err
}

// AddWalkdata insert a new Walkdata into database and returns
// last inserted Id on success.
func AddWalkdata(m *Walkdata) (id int64, err error) {
	// o := orm.NewOrm()
	// id, err = o.Insert(m)
	addWalkdataWithGeneralTS(m)
	id, err = 0, nil
	return id, err
}

func getWalkdataByUIdWithGeneralTS(aliuid string, datestamp int64) (v *Walkdata, err error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn(PKALIUID, aliuid)
	putPk.AddPrimaryKeyColumn(PKDATESTAMP, datestamp)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = walkDataTableName
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1

	getResp, err := tsclient.GetRow(getRowRequest)
	fmt.Println("getResp", getResp)

	colmap := getResp.GetColumnMap()
	fmt.Println("colmap len", len(colmap.Columns))
	wd := Walkdata{}
	if len(colmap.Columns) > 0 {
		step, _ := colmap.Columns[COLSTEP][0].Value.(float64)
		energy, _ := colmap.Columns[COLENERGY][0].Value.(float64)
		distance, _ := colmap.Columns[COLDISTANCE][0].Value.(float64)
		duration, _ := colmap.Columns[COLDURATION][0].Value.(float64)
		timestamp, _ := colmap.Columns[COLTIMESTAMP][0].Value.(int64)
		wd.Aliuid = aliuid
		wd.Step = step
		wd.Energy = energy
		wd.Distance = distance
		wd.Duration = duration
		wd.Timestamp = timestamp
		v = &wd
		return v, nil
	} else {
		return nil, nil
	}
}

// GetWalkdataById retrieves Walkdata by aliuid. Returns error if
// Id doesn't exist
func GetWalkdataByUId(aliuid string, datestamp int64) (v *Walkdata, err error) {
	v, err = getWalkdataByUIdWithGeneralTS(aliuid, datestamp)
	return v, err
}

func getBatchWalkdataByUId() {

}

func GetBatchWalkdataByUId(aliuid string, startdatestamp int64) (v []*Walkdata, err error) {
	return nil, nil
}

// // GetAllWalkdata retrieves all Walkdata matches certain condition. Returns empty list if
// // no records exist
// func GetAllWalkdata(query map[string]string, fields []string, sortby []string, order []string,
// 	offset int64, limit int64) (ml []interface{}, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(Walkdata))
// 	// query k=v
// 	for k, v := range query {
// 		// rewrite dot-notation to Object__Attribute
// 		k = strings.Replace(k, ".", "__", -1)
// 		qs = qs.Filter(k, v)
// 	}
// 	// order by:
// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			// 1) for each sort field, there is an associated order
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			// 2) there is exactly one order, all the sorted fields will be sorted by this order
// 			for _, v := range sortby {
// 				orderby := ""
// 				if order[0] == "desc" {
// 					orderby = "-" + v
// 				} else if order[0] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	var l []Walkdata
// 	qs = qs.OrderBy(sortFields...).RelatedSel()
// 	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
// 		if len(fields) == 0 {
// 			for _, v := range l {
// 				ml = append(ml, v)
// 			}
// 		} else {
// 			// trim unused fields
// 			for _, v := range l {
// 				m := make(map[string]interface{})
// 				val := reflect.ValueOf(v)
// 				for _, fname := range fields {
// 					m[fname] = val.FieldByName(fname).Interface()
// 				}
// 				ml = append(ml, m)
// 			}
// 		}
// 		return ml, nil
// 	}
// 	return nil, err
// }

// // UpdateWalkdata updates Walkdata by Id and returns error if
// // the record to be updated doesn't exist
// func UpdateWalkdataById(m *Walkdata) (err error) {
// 	o := orm.NewOrm()
// 	v := Walkdata{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// // DeleteWalkdata deletes Walkdata by Id and returns error if
// // the record to be deleted doesn't exist
// func DeleteWalkdata(id int64) (err error) {
// 	o := orm.NewOrm()
// 	v := Walkdata{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Walkdata{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
