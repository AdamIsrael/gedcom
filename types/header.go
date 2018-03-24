package types

// 1 CHAR UTF-8
//      2 VERS <VERSION_NUMBER>
// 1 LANG <LANGUAGE_OF_TEXT>
// 1 SOUR Ancestry.com Family Trees
//      2 VERS (2010.3)
//      2 NAME Ancestry.com Family Trees
//      2 CORP Ancestry.com
//          3 <<ADDRESS>>
//      2 DATA <<name of source>>
//          3 DATE <<publication date>>
//          3 COPR <<copyright source data>>
//              4 CONT|CONC <<copyright source data>>
// 1 DEST <<receiving system name>
// 1 DATE <<transmission date>>
//      2 TIME <<transmission time>>
// 1 SUBM @<XREF:SUBM>@
// 1 SUBN @<XREF:SUBN>@
// 1 FILE <FILE_NAME>
// 1 COPR <COPYRIGHT_GEDCOM_FILE>
// 1 GEDC
//      2 VERS 5.5
//      2 FORM LINEAGE-LINKED
// 1 PLAC
//      2 FORM <PLACE_HIERARCHY>
// 1 NOTE <GEDCOM_CONTENT_DESCRIPTION>
//      2 CONT|CONC <GEDCOM_CONTENT_DESCRIPTION>

type Header struct {
	ID              string // APPROVED_SYSTEM_ID - {Size=1:20}
	Version         string
	ProductName     string
	BusinessName    string
	BusinessAddress Address
	Language        string
}
