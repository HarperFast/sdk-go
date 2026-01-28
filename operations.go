package harper

const (
	OP_ADD_COMPONENT                   = "add_component"
	OP_ADD_CUSTOM_FUNCTION_PROJECT     = "add_custom_function_project"
	OP_ADD_NODE                        = "add_node"
	OP_ADD_ROLE                        = "add_role"
	OP_ADD_USER                        = "add_user"
	OP_ALTER_ROLE                      = "alter_role"
	OP_ALTER_USER                      = "alter_user"
	OP_CLUSTER_SET_ROUTES              = "cluster_set_routes"
	OP_CLUSTER_GET_ROUTES              = "cluster_get_routes"
	OP_CLUSTER_DELETE_ROUTES           = "cluster_delete_routes"
	OP_CLUSTER_NETWORK                 = "cluster_network"
	OP_CLUSTER_STATUS                  = "cluster_status"
	OP_CONFIGURE_CLUSTER               = "configure_cluster"
	OP_CREATE_ATTRIBUTE                = "create_attribute"
	OP_CREATE_AUTHENTICATION_TOKENS    = "create_authentication_tokens"
	OP_CREATE_DATABASE                 = "create_database"
	OP_CREATE_SCHEMA                   = "create_schema"
	OP_CREATE_TABLE                    = "create_table"
	OP_CSV_DATA_LOAD                   = "csv_data_load"
	OP_CSV_FILE_LOAD                   = "csv_file_load"
	OP_CSV_URL_LOAD                    = "csv_url_load"
	OP_CUSTOM_FUNCTIONS_STATUS         = "custom_functions_status"
	OP_DELETE_FILES_BEFORE             = "delete_files_before"
	OP_DELETE_TRANSACTION_LOG          = "delete_transaction_logs_before"
	OP_DEPLOY_COMPONENT                = "deploy_component"
	OP_DESCRIBE_ALL                    = "describe_all"
	OP_DESCRIBE_SCHEMA                 = "describe_schema"
	OP_DESCRIBE_DATABASE               = "describe_database"
	OP_DESCRIBE_TABLE                  = "describe_table"
	OP_DESCRIBE_METRIC                 = "describe_metric"
	OP_DELETE                          = "delete"
	OP_DELETE_AUDIT_LOGS_BEFORE        = "delete_audit_logs_before"
	OP_DELETE_RECORDS_BEFORE           = "delete_records_before"
	OP_DROP_ATTRIBUTE                  = "drop_attribute"
	OP_DROP_COMPONENT                  = "drop_component"
	OP_DROP_CUSTOM_FUNCTION            = "drop_custom_function"
	OP_DROP_CUSTOM_FUNCTION_PROJECT    = "drop_custom_function_project"
	OP_DROP_DATABASE                   = "drop_database"
	OP_DROP_ROLE                       = "drop_role"
	OP_DROP_SCHEMA                     = "drop_schema"
	OP_DROP_TABLE                      = "drop_table"
	OP_DROP_USER                       = "drop_user"
	OP_EXPORT_LOCAL                    = "export_local"
	OP_EXPORT_TO_S3                    = "export_to_s3"
	OP_GET_ANALYTICS                   = "get_analytics"
	OP_GET_BACKUP                      = "get_backup"
	OP_GET_CONFIGURATION               = "get_configuration"
	OP_GET_COMPONENT_FILE              = "get_component_file"
	OP_GET_COMPONENTS                  = "get_components"
	OP_GET_CUSTOM_FUNCTION             = "get_custom_function"
	OP_GET_CUSTOM_FUNCTIONS            = "get_custom_functions"
	OP_GET_FINGERPRINT                 = "get_fingerprint"
	OP_GET_JOB                         = "get_job"
	OP_IMPORT_FROM_S3                  = "import_from_s3"
	OP_INSERT                          = "insert"
	OP_INSTALL_NODE_MODULES            = "install_node_modules"
	OP_LIST_METRICS                    = "list_metrics"
	OP_LIST_ROLES                      = "list_roles"
	OP_LIST_USERS                      = "list_users"
	OP_PACKAGE_COMPONENT               = "package_component"
	OP_PACKAGE_CUSTOM_FUNCTION_PROJECT = "package_custom_function_project"
	OP_PURGE_STREAM                    = "purge_stream"
	OP_READ_AUDIT_LOG                  = "read_audit_log"
	OP_READ_LOG                        = "read_log"
	OP_READ_TRANSACTION_LOG            = "read_transaction_log"
	OP_REFRESH_OPERATION_TOKEN         = "refresh_operation_token"
	OP_REGISTRATION_INFO               = "registration_info"
	OP_REMOVE_NODE                     = "remove_node"
	OP_RESTART                         = "restart"
	OP_RESTART_SERVICE                 = "restart_service"
	OP_SEARCH_BY_CONDITIONS            = "search_by_conditions"
	OP_SEARCH_BY_HASH                  = "search_by_hash"
	OP_SEARCH_BY_ID                    = "search_by_id"
	OP_SEARCH_BY_VALUE                 = "search_by_value"
	OP_SEARCH_JOBS                     = "search_jobs_by_start_date"
	OP_SET_CONFIGURATION               = "set_configuration"
	OP_SET_COMPONENT_FILE              = "set_component_file"
	OP_SET_CUSTOM_FUNCTION             = "set_custom_function"
	OP_SET_LICENSE                     = "set_license"
	OP_SET_NODE_REPLICATION            = "set_node_replication"
	OP_SQL                             = "sql"
	OP_SYSTEM_INFORMATION              = "system_information"
	OP_UPDATE                          = "update"
	OP_UPDATE_NODE                     = "update_node"
	OP_UPSERT                          = "upsert"
	OP_USER_INFO                       = "user_info"
)

type Operation interface {
	Prepare() interface{}
}

type OpAddRole struct {
	Permission Permission `json:"permission"`
	Role       string     `json:"role"`
}

func (o OpAddRole) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAddRole
	}
	return Return{
		Operation: OP_ADD_ROLE,
		OpAddRole: o,
	}
}

type OpAlterRole struct {
	ID         string     `json:"id"`
	Role       string     `json:"role"`
	Permission Permission `json:"permission"`
}

func (o OpAlterRole) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAlterRole
	}
	return Return{
		Operation:   OP_ALTER_ROLE,
		OpAlterRole: o,
	}
}

// Describe Schema
type OpDescribeSchema struct {
	Schema string `json:"schema"`
}

func (o OpDescribeSchema) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpDescribeSchema
	}
	return Return{
		Operation:        OP_DESCRIBE_SCHEMA,
		OpDescribeSchema: o,
	}
}

// Set Routes
type OpSetRoutes struct {
	Server string  `json:"server"` // Must be either "hub" or "leaf"
	Routes []Route `json:"routes"`
}

func (o OpSetRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpSetRoutes
	}

	return Return{
		Operation:   OP_CLUSTER_SET_ROUTES,
		OpSetRoutes: o,
	}
}

// Delete Routes
type OpDeleteRoutes struct {
	Routes []Route `json:"routes"`
}

func (o OpDeleteRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpDeleteRoutes
	}

	return Return{
		Operation:      OP_CLUSTER_DELETE_ROUTES,
		OpDeleteRoutes: o,
	}
}

// Get Routes
type OpGetRoutes struct{}

func (o OpGetRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
	}

	return Return{
		Operation: OP_CLUSTER_GET_ROUTES,
	}
}

// Add Node
type OpAddNode struct {
	NodeName      string         `json:"node_name"`
	Host          string         `json:"host"`
	Port          int            `json:"port"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (o OpAddNode) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAddNode
	}

	return Return{
		Operation: OP_ADD_NODE,
		OpAddNode: o,
	}
}

// Remove Node
type OpRemoveNode struct {
	NodeName string `json:"node_name"`
}

func (o OpRemoveNode) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpRemoveNode
	}

	return Return{
		Operation:    OP_REMOVE_NODE,
		OpRemoveNode: o,
	}
}

// TODO: Figure out a better approach than having a superset of all ops params
// in one struct
type operation struct {
	Action              string                       `json:"action,omitempty"`
	Active              *bool                        `json:"active,omitempty"`
	Attribute           Attribute                    `json:"attribute,omitempty"`
	Attributes          []string                     `json:"attributes,omitempty"`
	CoalesceTime        bool                         `json:"coalesce_time,omitempty"`
	Conditions          []SearchCondition            `json:"conditions,omitempty"`
	Connections         []ConfigureClusterConnection `json:"connections,omitempty"`
	Company             string                       `json:"company,omitempty"`
	Comparator          string                       `json:"comparator,omitempty"`
	CustomMetricsWindow int64                        `json:"custom_metrics_window,omitempty"`
	CSVURL              string                       `json:"csv_url,omitempty"`
	Data                string                       `json:"data,omitempty"`
	Date                string                       `json:"date,omitempty"`
	Database            string                       `json:"database,omitempty"`
	DryRun              bool                         `json:"dry_run,omitempty"`
	EndTime             int64                        `json:"end_time,omitempty"`
	File                string                       `json:"file,omitempty"`
	FilePath            string                       `json:"file_path,omitempty"`
	Format              string                       `json:"format,omitempty"`
	From                string                       `json:"from,omitempty"`
	FromDate            string                       `json:"from_date,omitempty"`
	FunctionContent     string                       `json:"function_content,omitempty"`
	GetAttributes       AttributeList                `json:"get_attributes,omitempty"`
	HashAttribute       string                       `json:"hash_attribute,omitempty"`
	HashValues          interface{}                  `json:"hash_values,omitempty"`
	Host                string                       `json:"host,omitempty"`
	ID                  string                       `json:"id,omitempty"`
	IDs                 interface{}                  `json:"ids,omitempty"`
	Key                 string                       `json:"key,omitempty"`
	Limit               int                          `json:"limit,omitempty"`
	Metric              string                       `json:"metric,omitempty"`
	MetricTypes         []string                     `json:"metric_types,omitempty"`
	Name                string                       `json:"name,omitempty"`
	NodeName            string                       `json:"node_name,omitempty"`
	Offset              int                          `json:"offset,omitempty"`
	Operation           string                       `json:"operation"`
	Operator            string                       `json:"operator,omitempty"`
	Options             *PurgeStreamOptions          `json:"options,omitempty"`
	Order               string                       `json:"order,omitempty"`
	Package             string                       `json:"package,omitempty"`
	Password            string                       `json:"password,omitempty"`
	Path                string                       `json:"path,omitempty"`
	Payload             string                       `json:"payload,omitempty"`
	Permission          Permission                   `json:"permission,omitempty"`
	Port                int                          `json:"port,omitempty"`
	Project             string                       `json:"project,omitempty"`
	Projects            []string                     `json:"projects,omitempty"`
	Records             interface{}                  `json:"records,omitempty"`
	RefreshToken        string                       `json:"refresh_token,omitempty"`
	Role                string                       `json:"role,omitempty"`
	S3                  *S3Credentials               `json:"s3,omitempty"`
	Schema              string                       `json:"schema,omitempty"`
	SearchOperation     *SearchOperation             `json:"search_operation,omitempty"`
	SearchValues        interface{}                  `json:"search_values,omitempty"`
	Service             string                       `json:"service,omitempty"`
	SkipNodeModules     bool                         `json:"skip_node_modules,omitempty"`
	Sort                *Sort                        `json:"sort,omitempty"`
	Start               int                          `json:"start,omitempty"`
	StartTime           int64                        `json:"start_time,omitempty"`
	Subscriptions       []Subscription               `json:"subscriptions,omitempty"`
	SQL                 string                       `json:"sql,omitempty"`
	Table               string                       `json:"table,omitempty"`
	Timestamp           int64                        `json:"timestamp,omitempty"`
	ToDate              string                       `json:"to_date,omitempty"`
	Type                string                       `json:"type,omitempty"`
	Until               string                       `json:"until,omitempty"`
	Username            string                       `json:"username,omitempty"`
	Value               interface{}                  `json:"value,omitempty"`
}

func (o operation) Prepare() interface{} {
	return o
}
