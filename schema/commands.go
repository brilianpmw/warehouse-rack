package schema

type (
	//CMDType holds the type of commands
	CMDType = string
)

const (
	CMDCreateRack                       CMDType = "create_warehouse_rack"
	CMDRackIn                           CMDType = "rack"
	CMDStatus                           CMDType = "status"
	CMDExit                             CMDType = "exit"
	CMDSkuNumbersForProductWithExpDate  CMDType = "sku_numbers_for_product_with_exp_date"
	CMDRackOut                          CMDType = "rack_out"
	CMDSlotNumbersForProductWithExpDate CMDType = "slot_numbers_for_product_with_exp_date"
	CMDSlotNumberBySkuNumber            CMDType = "slot_number_for_sku_number"
)

// ValidCommandsByName holds the valid commands map
var ValidCommandsByName = map[string]bool{
	string(CMDCreateRack):                       true,
	string(CMDRackIn):                           true,
	string(CMDStatus):                           true,
	string(CMDExit):                             true,
	string(CMDSkuNumbersForProductWithExpDate):  true,
	string(CMDSlotNumbersForProductWithExpDate): true,
	string(CMDRackOut):                          true,
	string(CMDSlotNumberBySkuNumber):            true,
}

// CMDArgumentLength holds the exact arguments length to read for commands
var CMDArgumentLength = map[string]int{
	string(CMDCreateRack):                       1,
	string(CMDRackIn):                           2,
	string(CMDStatus):                           0,
	string(CMDExit):                             0,
	string(CMDRackOut):                          1,
	string(CMDSkuNumbersForProductWithExpDate):  1,
	string(CMDSlotNumbersForProductWithExpDate): 1,
	string(CMDSlotNumberBySkuNumber):            1,
}
