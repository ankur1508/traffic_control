// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Asn struct {
	Id          int64     `db:"id" json:"id"`
	Asn         int64     `db:"asn" json:"asn"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
	Links       AsnLinks  `json:"_links" db:-`
}

type AsnLinks struct {
	Self           string         `db:"self" json:"_self"`
	CachegroupLink CachegroupLink `json:"cachegroup" db:-`
}

// @Title getAsnById
// @Description retrieves the asn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn/{id} [get]
func getAsnById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Asn{}
	arg := Asn{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "asn/', id) as self "
	queryStr += ", concat('" + API_PATH + "cachegroup/', cachegroup) as cachegroup_id_ref"
	queryStr += " from asn where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getAsns
// @Description retrieves the asn
// @Accept  application/json
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn [get]
func getAsns(db *sqlx.DB) (interface{}, error) {
	ret := []Asn{}
	queryStr := "select *, concat('" + API_PATH + "asn/', id) as self "
	queryStr += ", concat('" + API_PATH + "cachegroup/', cachegroup) as cachegroup_id_ref"
	queryStr += " from asn"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postAsn
// @Description enter a new asn
// @Accept  application/json
// @Param                 Body body     Asn   true "Asn object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/asn [post]
func postAsn(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO asn("
	sqlString += "asn"
	sqlString += ",cachegroup"
	sqlString += ") VALUES ("
	sqlString += ":asn"
	sqlString += ",:cachegroup"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putAsn
// @Description modify an existing asnentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Asn   true "Asn object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/asn/{id}  [put]
func putAsn(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE asn SET "
	sqlString += "asn = :asn"
	sqlString += ",cachegroup = :cachegroup"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delAsnById
// @Description deletes asn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Asn
// @Resource /api/2.0
// @Router /api/2.0/asn/{id} [delete]
func delAsn(id int, db *sqlx.DB) (interface{}, error) {
	arg := Asn{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM asn WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
