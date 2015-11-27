// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateOpClassStmt struct {
	Opclassname  []Node    `json:"opclassname"`  /* qualified name (list of Value strings) */
	Opfamilyname []Node    `json:"opfamilyname"` /* qualified name (ditto); NIL if omitted */
	Amname       *string   `json:"amname"`       /* name of index AM opclass is for */
	Datatype     *TypeName `json:"datatype"`     /* datatype of indexed column */
	Items        []Node    `json:"items"`        /* List of CreateOpClassItem nodes */
	IsDefault    bool      `json:"isDefault"`    /* Should be marked as default for type? */
}

func (node CreateOpClassStmt) MarshalJSON() ([]byte, error) {
	type CreateOpClassStmtMarshalAlias CreateOpClassStmt
	return json.Marshal(map[string]interface{}{
		"CREATEOPCLASSSTMT": (*CreateOpClassStmtMarshalAlias)(&node),
	})
}

func (node *CreateOpClassStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["opclassname"] != nil {
		node.Opclassname, err = UnmarshalNodeArrayJSON(fields["opclassname"])
		if err != nil {
			return
		}
	}

	if fields["opfamilyname"] != nil {
		node.Opfamilyname, err = UnmarshalNodeArrayJSON(fields["opfamilyname"])
		if err != nil {
			return
		}
	}

	if fields["amname"] != nil {
		err = json.Unmarshal(fields["amname"], &node.Amname)
		if err != nil {
			return
		}
	}

	if fields["datatype"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["datatype"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(TypeName)
			node.Datatype = &val
		}
	}

	if fields["items"] != nil {
		node.Items, err = UnmarshalNodeArrayJSON(fields["items"])
		if err != nil {
			return
		}
	}

	if fields["isDefault"] != nil {
		err = json.Unmarshal(fields["isDefault"], &node.IsDefault)
		if err != nil {
			return
		}
	}

	return
}