package toolDetailsModel

type Tool struct {
		Name string `json:"Name"`		
}

type ToolDetails struct {
	Order_Id string `json:"Order_Id"`
	Created_On string `json:"Created_On"`
	Updated_On string `json:"Updated_On"`
	Tools []Tool `json:"Tools"`
}