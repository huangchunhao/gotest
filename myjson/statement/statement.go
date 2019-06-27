package statement

type(
	DataSource struct{
		Comment string
		DataSourceType string
		Num string
		Name string
		Ip string
		Port string
		Db string
		Type string
		UserName string
		PassWord string
	}

	Statement struct{
		DataSource DataSource
		InsertStatement string
		Result string
		DelStatement string
	}
)


