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
	PMessageMgmtMessages                   string                 `xml:"P_MESSAGEMGMT_MESSAGES,attr"`
	PArticleCraftProcess                   string                 `xml:"P_ARTICLE_CRAFT_PROCESS,attr"`
	PArticlePartNr                         string                 `xml:"P_ARTICLE_PARTNR,attr"`
	PArticleTypeNr                         string                 `xml:"P_ARTICLE_TYPENR,attr"`
	PArticleOrderNr                        string                 `xml:"P_ARTICLE_ORDERNR,attr"`
	PArticleDescr1                         string                 `xml:"P_ARTICLE_DESCR1,attr"`
	PArticleDescr2                         string                 `xml:"P_ARTICLE_DESCR2,attr"`
	PArticleManufacturer                   string                 `xml:"P_ARTICLE_MANUFACTURER,attr"`
	PArticleSupplier                       string                 `xml:"P_ARTICLE_SUPPLIER,attr"`
	PArticleNote                           string                 `xml:"P_ARTICLE_NOTE,attr"`
	PArticleHeight                         string                 `xml:"P_ARTICLE_HEIGHT,attr"`
	PArticleWidth                          string                 `xml:"P_ARTICLE_WIDTH,attr"`
	PArticleDepth                          string                 `xml:"P_ARTICLE_DEPTH,attr"`
	PArticleMountingSite                   string                 `xml:"P_ARTICLE_MOUNTINGSITE,attr"`
	PArticlePartType                       int                    `xml:"P_ARTICLE_PARTTYPE,attr"`
	PArticleProductSubGroup                int                    `xml:"P_ARTICLE_PRODUCTSUBGROUP,attr"`
	PArticleProductGroup                   int                    `xml:"P_ARTICLE_PRODUCTGROUP,attr"`
	PArticleQuantityUnit                   string                 `xml:"P_ARTICLE_QUANTITYUNIT,attr"`
	PArticlePriceUnit                      string                 `xml:"P_ARTICLE_PRICEUNIT,attr"`
	PArticlePictureFile                    string                 `xml:"P_ARTICLE_PICTUREFILE,attr"`
	PArticleWeight                         string                 `xml:"P_ARTICLE_WEIGHT,attr"`
	PArticleMountingSpace                  string                 `xml:"P_ARTICLE_MOUNTINGSPACE,attr"`
	PArticleIsAccessory                    string                 `xml:"P_ARTICLE_IS_ACCESSORY,attr"`
	PArticleErpNr                          string                 `xml:"P_ARTICLE_ERPNR,attr"`
	PArticleSalesPrice1                    string                 `xml:"P_ARTICLE_SALESPRICE_1,attr"`
	PArticleSalesPrice2                    string                 `xml:"P_ARTICLE_SALESPRICE_2,attr"`
	PArticlePurchasePrice1                 string                 `xml:"P_ARTICLE_PURCHASEPRICE_1,attr"`
	PArticlePurchasePrice2                 string                 `xml:"P_ARTICLE_PURCHASEPRICE_2,attr"`
	PArticlePackagingPrice1                string                 `xml:"P_ARTICLE_PACKAGINGPRICE_1,attr"`
	PArticlePackagingPrice2                string                 `xml:"P_ARTICLE_PACKAGINGPRICE_2,attr"`
	PArticleCertificateCe                  string                 `xml:"P_ARTICLE_CERTIFICATE_CE,attr"`
	PArticlePackagingQuantity              string                 `xml:"P_ARTICLE_PACKAGINGQUANTITY,attr"`
	PArticleCraftElectrical                string                 `xml:"P_ARTICLE_CRAFT_ELECTRICAL,attr"`
	PArticleCraftFluid                     string                 `xml:"P_ARTICLE_CRAFT_FLUID,attr"`
	PArticleCraftMechanics                 string                 `xml:"P_ARTICLE_CRAFT_MECHANICS,attr"`
	PArticleCraftHydraulics                string                 `xml:"P_ARTICLE_CRAFT_HYDRAULICS,attr"`
	PArticleCraftPneumatics                string                 `xml:"P_ARTICLE_CRAFT_PNEUMATICS,attr"`
	PArticleCraftLubrication               string                 `xml:"P_ARTICLE_CRAFT_LUBRICATION,attr"`
	PArticleCraftCooling                   string                 `xml:"P_ARTICLE_CRAFT_COOLING,attr"`
	PArticleProductTopGroup                string                 `xml:"P_ARTICLE_PRODUCTTOPGROUP,attr"`
	PArticleGroupSymbolMacro               string                 `xml:"P_ARTICLE_GROUPSYMBOLMACRO,attr"`
	PArticleExternalDoc1                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_1,attr"`
	PArticleExternalDoc2                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_2,attr"`
	PArticleExternalDoc3                   string                 `xml:"P_ARTICLE_EXTERNAL_DOCUMENT_3,attr"`
	PArticleSpacingLeft                    string                 `xml:"P_ARTICLE_SPACING_LEFT,attr"`
	PArticleSpacingRight                   string                 `xml:"P_ARTICLE_SPACING_RIGHT,attr"`
	PArticleSpacingAbove                   string                 `xml:"P_ARTICLE_SPACING_ABOVE,attr"`
	PArticleSpacingBelow                   string                 `xml:"P_ARTICLE_SPACING_BELOW,attr"`
	PArticleSpacingFront                   string                 `xml:"P_ARTICLE_SPACING_FRONT,attr"`
	PArticleSpacingRear                    string                 `xml:"P_ARTICLE_SPACING_REAR,attr"`
	PArticleSnapHeight                     string                 `xml:"P_ARTICLE_SNAPHEIGHT,attr"`
	PArticleMiddleOffset                   string                 `xml:"P_ARTICLE_MIDDLEOFFSET,attr"`
	PArticleRefConstructionName            string                 `xml:"P_ARTICLE_REF_CONSTRUCTION_NAME,attr"`
	PArticleEcabinetMacro                  string                 `xml:"P_ARTICLE_ECABINET_MACRO,attr"`
	PArticleExternalPlacement              string                 `xml:"P_ARTICLE_EXTERNAL_PLACEMENT,attr"`
	PArticleDiscount                       string                 `xml:"P_ARTICLE_DISCOUNT,attr"`
	PArticleCanBeLinedUp                   string                 `xml:"P_ARTICLE_CAN_BE_LINED_UP,attr"`
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
	PArticleConnectionWireCrossSectionUnit string                 `xml:"P_ARTICLE_CONNECTION_WIRECROSSSECTION_UNIT,attr"`
	PArticleDiscontinued                   string                 `xml:"P_ARTICLE_DISCONTINUED,attr"`
	PArticleCraftCoolingLubricant          string                 `xml:"P_ARTICLE_CRAFT_COOLINGLUBRICANT,attr"`
	PArticleCraftGasTechnology             string                 `xml:"P_ARTICLE_CRAFT_GASTECHNOLOGY,attr"`
	PArticleCraftFluidUndefined            string                 `xml:"P_ARTICLE_CRAFT_FLUID_UNDEFINED,attr"`
	PArticleInstallationDepth              string                 `xml:"P_ARTICLE_INSTALLATION_DEPTH,attr"`
	PArticleEdpChecksum                    string                 `xml:"P_ARTICLE_EDP_CHECKSUM,attr"`
	PArticleRefTerminalOffsetX             string                 `xml:"P_ARTICLE_REF_TERMINAL_OFFSET_X,attr"`
	PArticleRefTerminalOffsetY             string                 `xml:"P_ARTICLE_REF_TERMINAL_OFFSET_Y,attr"`
	PArticleDisassembleMode                string                 `xml:"P_ARTICLE_DISASSEMBLE_MODE,attr"`
	PPartLastChange                        string                 `xml:"P_PART_LASTCHANGE,attr"`
	PPartCreate                            string                 `xml:"P_PART_CREATE,attr"`
	PArticleRefTerminalName                string                 `xml:"P_ARTICLE_REF_TERMINAL_NAME,attr"`
	PPartCreateDateUtc                     string                 `xml:"P_PART_CREATE_DATE_UTC,attr"`
	PPartLastChangeDateUtc                 string                 `xml:"P_PART_LASTCHANGE_DATE_UTC,attr"`
	FreeProperties                         []FreeProperty         `xml:"freeproperty"`
	AttributePositions                     []AttributePosition    `xml:"attributeposition"`
	AccessoryPositions                     []AccessoryPosition    `xml:"accessoryposition"`
	ConstructionPositions                  []ConstructionPosition `xml:"constructionPosition"`
	Variants                               []Variant              `xml:"variant"`
}

type FreeProperty struct {
	PArticleFreeDataDescription string `xml:"P_ARTICLE_FREE_DATA_DESCRIPTION,attr"`
	Pos                         string `xml:"pos,attr"`
	PArticleFreeDataValue       string `xml:"P_ARTICLE_FREE_DATA_VALUE,attr,omitempty"`
}

type AttributePosition struct {
	Pos                    string `xml:"pos,attr"`
	PArticleAttributeValue string `xml:"P_ARTICLE_ATTRIBUTE_VALUE,attr"`
}

type AccessoryPosition struct {
	Necessary string `xml:"necessary,attr"`
	PartNr    string `xml:"partnr,attr"`
	PartType  string `xml:"parttype,attr"`
	Pos       string `xml:"pos,attr"`
}

type ConstructionPosition struct {
	OffsetX string `xml:"offsetx,attr"`
	OffsetY string `xml:"offsety,attr"`
	Pos     string `xml:"pos,attr"`
	Name    string `xml:"name,attr,omitempty"`
}

type Variant struct {
	PArticleCharacteristics           string             `xml:"P_ARTICLE_CHARACTERISTICS,attr"`
	PArticlePlcIsBusCoupler           string             `xml:"P_ARTICLE_PLCISBUSCOUPLER,attr"`
	PArticlePlcIsCpu                  string             `xml:"P_ARTICLE_PLCISCPU,attr"`
	PArticleVariant                   string             `xml:"P_ARTICLE_VARIANT,attr"`
	PArticleAssemblyPosPlaceSpreading string             `xml:"P_ARTICLE_ASSEMBLY_POS_PLACE_SPREADED,attr"`
	PArticlePlcIsPowerSupply          string             `xml:"P_ARTICLE_PLCISPOWERSUPPLY,attr"`
	PArticlePlcIsBusDistributor       string             `xml:"P_ARTICLE_PLCISBUSDISTRIBUTOR,attr"`
	PArticleCableLength               string             `xml:"P_ARTICLE_CABLELENGTH,attr"`
	PArticleElectricalPower           string             `xml:"P_ARTICLE_ELECTRICALPOWER,attr"`
	PArticlePowerDissipation          string             `xml:"P_ARTICLE_POWERDISSIPATION,attr"`
	PArticlePanelMountingSpace        string             `xml:"P_ARTICLE_PANELMOUNTINGSPACE,attr"`
	PArticleDoorMountingSpace         string             `xml:"P_ARTICLE_DOORMOUNTINGSPACE,attr"`
	PArticleAddressRange              string             `xml:"P_ARTICLE_ADDRESSRANGE,attr"`
	PArticleIntrinsicSafety           string             `xml:"P_ARTICLE_INTRINSICSAFETY,attr"`
	PArticleShortCircuitResistant     string             `xml:"P_ARTICLE_SHORTCIRCUITRESISTANT,attr"`
	PArticlePanelHeight               string             `xml:"P_ARTICLE_PANELHEIGHT,attr"`
	PArticlePanelWidth                string             `xml:"P_ARTICLE_PANELWIDTH,attr"`
	PArticlePanelDepth                string             `xml:"P_ARTICLE_PANELDEPTH,attr"`
	PArticleDoorHeight                string             `xml:"P_ARTICLE_DOORHEIGHT,attr"`
	PArticleDoorWidth                 string             `xml:"P_ARTICLE_DOORWIDTH,attr"`
	PArticleDoorDepth                 string             `xml:"P_ARTICLE_DOORDEPTH,attr"`
	PArticlePressure                  string             `xml:"P_ARTICLE_PRESSURE,attr"`
	PArticleAdjustRange               string             `xml:"P_ARTICLE_ADJUSTRANGE,attr"`
	PArticleFlow                      string             `xml:"P_ARTICLE_FLOW,attr"`
	PArticleWireCrossSectionUnit      string             `xml:"P_ARTICLE_WIRECROSSSECTION_UNIT,attr"`
	PArticleModulePosPlaceSpreading   string             `xml:"P_ARTICLE_MODULE_POS_PLACE_SPREADED,attr"`
	PArticlePlcIsMountedOnHeadModule  string             `xml:"P_ARTICLE_PLCISMOUNTEDONHEADMODULE,attr"`
	FunctionTemplates                 []FunctionTemplate `xml:"functiontemplate"`
}

type FunctionTemplate struct {
	ConnectionDesignation string `xml:"connectionDesignation,attr,omitempty"`
	ConnectionDescription string `xml:"connectiondescription,attr,omitempty"`
	FunctionDefCategory   string `xml:"functiondefcategory,attr"`
	FunctionDefGroup      string `xml:"functiondefgroup,attr"`
	FunctionDefId         string `xml:"functiondefid,attr"`
	HasLed                string `xml:"hasled,attr"`
	HasPlugAdapter        string `xml:"hasplugadapter,attr"`
	IntrinsicSafety       string `xml:"intrinsicsafety,attr"`
	ManualModuleTemplate  string `xml:"manualmoduletemplate,attr"`
	PlcBusSystem          string `xml:"plcbussystem,attr"`
	Pos                   string `xml:"pos,attr"`
	SafetyRelevant        string `xml:"safetyrelevant,attr"`
	Symbol                string `xml:"symbol,attr"`
	TerminalFunction      string `xml:"terminalfunction,attr"`
	IndexStartAddress     string `xml:"indexstartaddress,attr"`
}
