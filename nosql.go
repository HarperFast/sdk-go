package harper

type AffectedResponse struct {
	MessageResponse
	SkippedHashes  []interface{} `json:"skipped_hashes"`
	InsertedHashes []interface{} `json:"inserted_hashes"`
	UpdatedHashes  []interface{} `json:"update_hashes"` // (sic) not updated_hashes
	DeletedHashes  []interface{} `json:"deleted_hashes"`
	UpsertedHashes []interface{} `json:"upserted_hashes"`
	// returned hashes can be of any JSON primitive
}

type SearchByConditionsOptions struct {
	Operator string `json:"operator,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Sort     Sort   `json:"sort,omitempty"`
}

type SearchCondition struct {
	Attribute  string             `json:"attribute,omitempty"`
	Comparator string             `json:"comparator,omitempty"`
	Value      interface{}        `json:"value,omitempty"`
	Operator   string             `json:"operator,omitempty"`
	Conditions []*SearchCondition `json:"conditions,omitempty"`
}

type SearchConditions []SearchCondition

type Sort struct {
	Attribute  string `json:"attribute,omitempty"`
	Descending bool   `json:"descending,omitempty"`
	Next       *Sort  `json:"next,omitempty"`
}

// Insert inserts one or more JSON objects into a table
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Insert(schema, table string, records interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_INSERT,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, &result)
	return &result, err
}

// Update updates one or more JSON objects in a table.
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Update(schema, table string, records interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_UPDATE,
		Schema:    schema,
		Table:     table,
		Records:   records,
	}, &result)
	return &result, err
}

// Update updates one or more JSON objects in a table.
// Hash value of the inserted JSON record MUST be present.
func (c *Client) Upsert(database, table string, records interface{}) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation: OP_UPSERT,
		Database:  database,
		Table:     table,
		Records:   records,
	}, &result)

	return &result, err
}

// Delete delete one or more JSON objects from a table.
// hashValues must be an array of slice
func (c *Client) Delete(schema, table string, hashValues AttributeList) (*AffectedResponse, error) {
	result := AffectedResponse{}
	err := c.opRequest(operation{
		Operation:  OP_DELETE,
		Schema:     schema,
		Table:      table,
		HashValues: hashValues,
	}, &result)
	return &result, err
}

// SearchByHash fetches records based on the table's hash field
// (i.e. primary key).
func (c *Client) SearchByHash(schema, table string, v interface{}, hashValues AttributeList, getAttributes AttributeList) error {
	return c.opRequest(operation{
		Operation:     OP_SEARCH_BY_HASH,
		Schema:        schema,
		Table:         table,
		HashValues:    hashValues,
		GetAttributes: getAttributes,
	}, &v)
}

func (c *Client) SearchById(database, table string, v interface{}, ids interface{}, getAttributes AttributeList) error {
	return c.opRequest(operation{
		Operation:     OP_SEARCH_BY_ID,
		Database:      database,
		Table:         table,
		IDs:           ids,
		GetAttributes: getAttributes,
	}, &v)
}

// SearchByValue fetches records based on the value of an attribute
// Wilcards are allowed in `searchValue`
func (c *Client) SearchByValue(schema, table string, v interface{}, attribute Attribute, value interface{}, getAttributes AttributeList) error {
	op := operation{
		Operation:     OP_SEARCH_BY_VALUE,
		Schema:        schema,
		Table:         table,
		Attribute:     attribute,
		Value:         value,
		GetAttributes: getAttributes,
	}
	return c.opRequest(op, &v)
}

func (c *Client) SearchByConditions(database, table string, v interface{}, conditions []SearchCondition, getAttributes AttributeList, options SearchByConditionsOptions) error {
	op := operation{
		Operation:     OP_SEARCH_BY_CONDITIONS,
		Database:      database,
		Table:         table,
		Conditions:    conditions,
		Offset:        options.Offset,
		Limit:         options.Limit,
		GetAttributes: getAttributes,
	}

	if (options.Sort != Sort{}) {
		op.Sort = &options.Sort
	}

	return c.opRequest(op, &v)
}
