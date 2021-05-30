package router

type Routes []Route

var routes = Routes{
	Route{
		"GetAllEmployeesDetails",
		"GET",
		"/employee",
		getEmployeesDetails,
	},
	Route{
		"GetEmployeeDetails",
		"GET",
		"/employee/{id}",
		getEmployeeDetails,
	},
	Route{
		"CreateEmployee",
		"POST",
		"/employee",
		createEmployee,
	},
	Route{
		"UpdateEmployeeDetails",
		"GET",
		"/employee",
		updateEmployeeDetails,
	},
	Route{
		"DeleteEmployee",
		"DELETE",
		"/employee/{id}",
		deleteEmployee,
	},
	Route{
		"GetCommonManager",
		"GET",
		"/commonmgr",
		getCommonManager,
	},
	Route{
		"changeManager",
		"PUT",
		"/changeManager",
		changeManager,
	},
	Route{
		"MoveEmployees",
		"PUT",
		"/moveEmployees",
		moveEmployees,
	},
}
