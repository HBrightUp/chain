CREATE TABLE IF NOT EXISTS user_attribute
(
    `UUID` 				UInt64,
    `Attribute_Id` 		UInt64,
    `Attribute_Name` 	String,
    `Attribute_type` 	UInt8,
    `Value_bit` 		Nullable(Int8),
    `Value_ubit` 		Nullable(UInt8),
    `Value_int64`	 	Nullable(Int64),
    `Value_uint64` 		Nullable(UInt64),
    `Value_double` 		Nullable(Float64),
    `Value_string` 		Nullable(String),
    `Value_date` 		Nullable(DateTime)
)
ENGINE = MergeTree
ORDER BY UUID
SETTINGS index_granularity = 8192;


CREATE TABLE IF NOT EXISTS user_action
(
    `UUID` 					UInt64,
    `Action_ID` 			UInt64,
    `Action_Name` 			String,
    `Action_Input_Type` 	UInt8,
    `Action_Output_Type` 	UInt8,
    `Value_bit` 			Nullable(Int8),
    `Value_ubit` 			Nullable(UInt8),
    `Value_int64`	 		Nullable(Int64),
    `Value_uint64` 			Nullable(UInt64),
    `Value_double` 			Nullable(Float64),
    `Value_string` 			Nullable(String),
    `Value_date` 			Nullable(DateTime)
)
ENGINE = MergeTree
ORDER BY UUID
SETTINGS index_granularity = 8192;

-- default.business definition

CREATE TABLE IF NOT EXISTS business
(

    `UUID`              UInt64,       --主键
    `Business_ID`       String,       --总业务类型
    `Business_Name`     String,           
    `SubBusiness_ID`    String,       --子业务类型
    `Logic_ID`          String,       --操作类型
    `Logic_Name`        String,       --
    `Fund_ID`           String,       --资金名称
    `Fund_Name`         String,       --
    `Account_ID`        UInt64,       --用户id
    `Transaction_ID`    String,       --第三方对应的交易流水id
    `Transaction_Name`  String,       --
    `Value_String`      String,       --交易金额1
    `Transaction_Date`  DateTime      --业务日期
)
ENGINE = MergeTree
ORDER BY UUID
SETTINGS index_granularity = 8192;