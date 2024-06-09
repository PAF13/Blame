package parts

import (
	"encoding/xml"
)

type PartsManagement struct {
	XMLName    xml.Name `xml:"partsmanagement"`
	Count      int      `xml:"count,attr"`
	LengthUnit string   `xml:"length-unit,attr"`
	WeightUnit string   `xml:"weight-unit,attr"`
	Type       string   `xml:"type,attr"`
	Build      string   `xml:"build,attr"`
	Version    string   `xml:"version,attr"`
	Parts      []Part   `xml:"part"`
}

type Part struct {
	PMessageMgmtMessages                   int                    `xml:"P_MESSAGEMGMT_MESSAGES,attr"`
	PArticleCraftProcess                   int                    `xml:"P_ARTICLE_CRAFT_PROCESS,attr"`
	PArticlePartNr                         string                 `xml:"P_ARTICLE_PARTNR,attr"`
	PArticleTypeNr                         string                 `xml:"P_ARTICLE_TYPENR,attr"`
	PArticleOrderNr                        string                 `xml:"P_ARTICLE_ORDERNR,attr"`
	PArticleDescr1                         string                 `xml:"P_ARTICLE_DESCR1,attr"`
	PArticleDescr2                         string                 `xml:"P_ARTICLE_DESCR2,attr"`
	PArticleManufacturer                   string                 `xml:"P_ARTICLE_MANUFACTURER,attr"`
	PArticleSupplier                       string                 `xml:"P_ARTICLE_SUPPLIER,attr"`
	PArticleNote                           string                 `xml:"P_ARTICLE_NOTE,attr"`
	PArticleHeight                         float64                `xml:"P_ARTICLE_HEIGHT,attr"`
	PArticleWidth                          float64                `xml:"P_ARTICLE_WIDTH,attr"`
	PArticleDepth                          float64                `xml:"P_ARTICLE_DEPTH,attr"`
	PArticleMountingSite                   int                    `xml:"P_ARTICLE_MOUNTINGSITE,attr"`
	PArticlePartType                       int                    `xml:"P_ARTICLE_PARTTYPE,attr"`
	PArticleProductSubGroup                int                    `xml:"P_ARTICLE_PRODUCTSUBGROUP,attr"`
	PArticleProductGroup                   int                    `xml:"P_ARTICLE_PRODUCTGROUP,attr"`
	PArticleQuantityUnit                   string                 `xml:"P_ARTICLE_QUANTITYUNIT,attr"`
	PArticlePriceUnit                      int                    `xml:"P_ARTICLE_PRICEUNIT,attr"`
	PArticlePictureFile                    string                 `xml:"P_ARTICLE_PICTUREFILE,attr"`
	PArticleWeight                         float64                `xml:"P_ARTICLE_WEIGHT,attr"`
	PArticleMountingSpace                  float64                `xml:"P_ARTICLE_MOUNTINGSPACE,attr"`
	PArticleIsAccessory                    int                    `xml:"P_ARTICLE_IS_ACCESSORY,attr"`
	PArticleErpNr                          string                 `xml:"P_ARTICLE_ERPNR,attr"`
	PArticleSalesPrice1                    float64                `xml:"P_ARTICLE_SALESPRICE_1,attr"`
	PArticleSalesPrice2                    float64                `xml:"P_ARTICLE_SALESPRICE_2,attr"`
	PArticlePurchasePrice1                 float64                `xml:"P_ARTICLE_PURCHASEPRICE_1,attr"`
	PArticlePurchasePrice2                 float64                `xml:"P_ARTICLE_PURCHASEPRICE_2,attr"`
	PArticlePackagingPrice1                float64                `xml:"P_ARTICLE_PACKAGINGPRICE_1,attr"`
	PArticlePackagingPrice2                float64                `xml:"P_ARTICLE_PACKAGINGPRICE_2,attr"`
	PArticleCertificateCe                  int                    `xml:"P_ARTICLE_CERTIFICATE_CE,attr"`
	PArticlePackagingQuantity              int                    `xml:"P_ARTICLE_PACKAGINGQUANTITY,attr"`
	PArticleCraftElectrical                int                    `xml:"P_ARTICLE_CRAFT_ELECTRICAL,attr"`
	PArticleCraftFluid                     int                    `xml:"P_ARTICLE_CRAFT_FLUID,attr"`
	PArticleCraftMechanics                 int                    `xml:"P_ARTICLE_CRAFT_MECHANICS,attr"`
	PArticleCraftHydraulics                int                    `xml:"P_ARTICLE_CRAFT_HYDRAULICS,attr"`
	PArticleCraftPneumatics                int                    `xml:"P_ARTICLE_CRAFT_PNEUMATICS,attr"`
	PArticleCraftLubrication               int                    `xml:"P_ARTICLE_CRAFT_LUBRICATION,attr"`
	PArticleCraftCooling                   int                    `xml:"P_ARTICLE_CRAFT_COOLING,attr"`
	PArticleProductTopGroup                int                    `xml:"P_ARTICLE_PRODUCTTOPGROUP,attr"`
	PArticleGroupSymbolMacro               string                 `xml:"P_ARTICLE_GROUPSYMBOLMACRO,attr"`
	PArticleExternalDoc1                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_1,attr"`
	PArticleExternalDoc2                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_2,attr"`
	PArticleExternalDoc3                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_3,attr"`
	PArticleSpacingLeft                    int                    `xml:"P_ARTICLE_SPACING_LEFT,attr"`
	PArticleSpacingRight                   int                    `xml:"P_ARTICLE_SPACING_RIGHT,attr"`
	PArticleSpacingAbove                   int                    `xml:"P_ARTICLE_SPACING_ABOVE,attr"`
	PArticleSpacingBelow                   int                    `xml:"P_ARTICLE_SPACING_BELOW,attr"`
	PArticleSpacingFront                   int                    `xml:"P_ARTICLE_SPACING_FRONT,attr"`
	PArticleSpacingRear                    int                    `xml:"P_ARTICLE_SPACING_REAR,attr"`
	PArticleSnapHeight                     int                    `xml:"P_ARTICLE_SNAPHEIGHT,attr"`
	PArticleMiddleOffset                   int                    `xml:"P_ARTICLE_MIDDLEOFFSET,attr"`
	PArticleRefConstructionName            string                 `xml:"P_ARTICLE_REF_CONSTRUCTION_NAME,attr"`
	PArticleEcabinetMacro                  string                 `xml:"P_ARTICLE_ECABINET_MACRO,attr"`
	PArticleExternalPlacement              int                    `xml:"P_ARTICLE_EXTERNAL_PLACEMENT,attr"`
	PArticleDiscount                       int                    `xml:"P_ARTICLE_DISCOUNT,attr"`
	PArticleCanBeLinedUp                   int                    `xml:"P_ARTICLE_CAN_BE_LINED_UP,attr"`
	PArticleExternalDoc4                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_4,attr"`
	PArticleExternalDoc5                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_5,attr"`
	PArticleExternalDoc6                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_6,attr"`
	PArticleExternalDoc7                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_7,attr"`
	PArticleExternalDoc8                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_8,attr"`
	PArticleExternalDoc10                  string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_10,attr"`
	PArticleExternalDoc11                  string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_11,attr"`
	PArticleExternalDoc12                  string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_12,attr"`
	PArticleExternalDoc13                  string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_13,attr"`
	PArticleExternalDoc20                  string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_20,attr"`
	PArticleConnectionWireCrossSectionUnit int                    `xml:"P_ARTICLE_CONNECTION_WIRECROSSSECTION_UNIT,attr"`
	PArticleDiscontinued                   int                    `xml:"P_ARTICLE_DISCONTINUED,attr"`
	PArticleCraftCoolingLubricant          int                    `xml:"P_ARTICLE_CRAFT_COOLINGLUBRICANT,attr"`
	PArticleCraftGasTechnology             int                    `xml:"P_ARTICLE_CRAFT_GASTECHNOLOGY,attr"`
	PArticleCraftFluidUndefined            int                    `xml:"P_ARTICLE_CRAFT_FLUID_UNDEFINED,attr"`
	PArticleInstallationDepth              int                    `xml:"P_ARTICLE_INSTALLATION_DEPTH,attr"`
	PArticleEdpChecksum                    string                 `xml:"P_ARTICLE_EDP_CHECKSUM,attr"`
	PArticleRefTerminalOffsetX             int                    `xml:"P_ARTICLE_REF_TERMINAL_OFFSET_X,attr"`
	PArticleRefTerminalOffsetY             int                    `xml:"P_ARTICLE_REF_TERMINAL_OFFSET_Y,attr"`
	PArticleDisassembleMode                int                    `xml:"P_ARTICLE_DISASSEMBLE_MODE,attr"`
	PPartLastChange                        string                 `xml:"P_PART_LASTCHANGE,attr"`
	PPartCreate                            string                 `xml:"P_PART_CREATE,attr"`
	PArticleRefTerminalName                string                 `xml:"P_ARTICLE_REF_TERMINAL_NAME,attr"`
	PPartCreateDateUtc                     int64                  `xml:"P_PART_CREATE_DATE_UTC,attr"`
	PPartLastChangeDateUtc                 int64                  `xml:"P_PART_LASTCHANGE_DATE_UTC,attr"`
	FreeProperties                         []FreeProperty         `xml:"freeproperty"`
	AttributePositions                     []AttributePosition    `xml:"attributeposition"`
	AccessoryPositions                     []AccessoryPosition    `xml:"accessoryposition"`
	ConstructionPositions                  []ConstructionPosition `xml:"constructionPosition"`
	Variants                               []Variant              `xml:"variant"`
}

type FreeProperty struct {
	PArticleFreeDataDescription string `xml:"P_ARTICLE_FREE_DATA_DESCRIPTION,attr"`
	Pos                         int    `xml:"pos,attr"`
	PArticleFreeDataValue       string `xml:"P_ARTICLE_FREE_DATA_VALUE,attr,omitempty"`
}

type AttributePosition struct {
	Pos                    int    `xml:"pos,attr"`
	PArticleAttributeValue string `xml:"P_ARTICLE_ATTRIBUTE_VALUE,attr"`
}

type AccessoryPosition struct {
	Necessary int    `xml:"necessary,attr"`
	PartNr    string `xml:"partnr,attr"`
	PartType  int    `xml:"parttype,attr"`
	Pos       int    `xml:"pos,attr"`
}

type ConstructionPosition struct {
	OffsetX int    `xml:"offsetx,attr"`
	OffsetY int    `xml:"offsety,attr"`
	Pos     int    `xml:"pos,attr"`
	Name    string `xml:"name,attr,omitempty"`
}

type Variant struct {
	PArticleCharacteristics           string             `xml:"P_ARTICLE_CHARACTERISTICS,attr"`
	PArticlePlcIsBusCoupler           int                `xml:"P_ARTICLE_PLCISBUSCOUPLER,attr"`
	PArticlePlcIsCpu                  int                `xml:"P_ARTICLE_PLCISCPU,attr"`
	PArticleVariant                   int                `xml:"P_ARTICLE_VARIANT,attr"`
	PArticleAssemblyPosPlaceSpreading int                `xml:"P_ARTICLE_ASSEMBLY_POS_PLACE_SPREADED,attr"`
	PArticlePlcIsPowerSupply          int                `xml:"P_ARTICLE_PLCISPOWERSUPPLY,attr"`
	PArticlePlcIsBusDistributor       int                `xml:"P_ARTICLE_PLCISBUSDISTRIBUTOR,attr"`
	PArticleCableLength               int                `xml:"P_ARTICLE_CABLELENGTH,attr"`
	PArticleElectricalPower           string             `xml:"P_ARTICLE_ELECTRICALPOWER,attr"`
	PArticlePowerDissipation          int                `xml:"P_ARTICLE_POWERDISSIPATION,attr"`
	PArticlePanelMountingSpace        int                `xml:"P_ARTICLE_PANELMOUNTINGSPACE,attr"`
	PArticleDoorMountingSpace         int                `xml:"P_ARTICLE_DOORMOUNTINGSPACE,attr"`
	PArticleAddressRange              string             `xml:"P_ARTICLE_ADDRESSRANGE,attr"`
	PArticleIntrinsicSafety           int                `xml:"P_ARTICLE_INTRINSICSAFETY,attr"`
	PArticleShortCircuitResistant     int                `xml:"P_ARTICLE_SHORTCIRCUITRESISTANT,attr"`
	PArticlePanelHeight               int                `xml:"P_ARTICLE_PANELHEIGHT,attr"`
	PArticlePanelWidth                int                `xml:"P_ARTICLE_PANELWIDTH,attr"`
	PArticlePanelDepth                int                `xml:"P_ARTICLE_PANELDEPTH,attr"`
	PArticleDoorHeight                int                `xml:"P_ARTICLE_DOORHEIGHT,attr"`
	PArticleDoorWidth                 int                `xml:"P_ARTICLE_DOORWIDTH,attr"`
	PArticleDoorDepth                 int                `xml:"P_ARTICLE_DOORDEPTH,attr"`
	PArticlePressure                  int                `xml:"P_ARTICLE_PRESSURE,attr"`
	PArticleAdjustRange               int                `xml:"P_ARTICLE_ADJUSTRANGE,attr"`
	PArticleFlow                      int                `xml:"P_ARTICLE_FLOW,attr"`
	PArticleWireCrossSectionUnit      int                `xml:"P_ARTICLE_WIRECROSSSECTION_UNIT,attr"`
	PArticleModulePosPlaceSpreading   int                `xml:"P_ARTICLE_MODULE_POS_PLACE_SPREADED,attr"`
	PArticlePlcIsMountedOnHeadModule  int                `xml:"P_ARTICLE_PLCISMOUNTEDONHEADMODULE,attr"`
	FunctionTemplates                 []FunctionTemplate `xml:"functiontemplate"`
}

type FunctionTemplate struct {
	ConnectionDesignation string `xml:"connectionDesignation,attr,omitempty"`
	ConnectionDescription string `xml:"connectiondescription,attr,omitempty"`
	FunctionDefCategory   int    `xml:"functiondefcategory,attr"`
	FunctionDefGroup      int    `xml:"functiondefgroup,attr"`
	FunctionDefId         int    `xml:"functiondefid,attr"`
	HasLed                int    `xml:"hasled,attr"`
	HasPlugAdapter        int    `xml:"hasplugadapter,attr"`
	IntrinsicSafety       int    `xml:"intrinsicsafety,attr"`
	ManualModuleTemplate  int    `xml:"manualmoduletemplate,attr"`
	PlcBusSystem          int    `xml:"plcbussystem,attr"`
	Pos                   int    `xml:"pos,attr"`
	SafetyRelevant        int    `xml:"safetyrelevant,attr"`
	Symbol                string `xml:"symbol,attr"`
	TerminalFunction      int    `xml:"terminalfunction,attr"`
	IndexStartAddress     int    `xml:"indexstartaddress,attr"`
}
