package main

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	///////str[output.csv-name]=[]string{original-name,type}
	h, _ := os.Getwd()
	currentPath := filepath.Join(h, "testmapping")

	csvFile, err := os.Open(filepath.Join(currentPath, "output (29).csv"))
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	row, err := csvReader.Read()
	fmt.Println(",", row)

	/////////////////////////////////////////////////
	/*file, err := os.Open(filepath.Join(currentPath, "incidentMapping.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fieldsMappingBytes, err := ioutil.ReadAll(file)
	fieldsMapping := string(fieldsMappingBytes)
	fieldsMappingSplits := strings.Split(fieldsMapping, ",")
	for _, valMap := range fieldsMappingSplits {

		//make the field names lower and remove white spaces
		valMap = strings.ToLower(valMap)
		valMap = strings.Replace(valMap, " ", "", -1)
		valMap = "\"" + valMap + "\" : {\"type\" : \"keyword\"},"
		fmt.Println(valMap)
	}*/
	/////////////////////////////////////////////////

	/*
		file, err := os.Open(filepath.Join(currentPath, "fieldsMapping.txt"))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//declaration to exclude out all regex
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}

		fieldsMappingBytes, err := ioutil.ReadAll(file)
		fieldsMapping := string(fieldsMappingBytes)
		fieldsMappingSplits := strings.Split(fieldsMapping, ",")
		finalMapping := make(map[string][]string)
		for key, val := range fieldsMappingSplits {
			dataType := "keyword"

			//get rid of -index from the data
			if idx := strings.Index(val, "-index"); idx != -1 {
				fieldsMappingSplits[key] = val[:idx]
			}

			//make the field names lower and remove white spaces
			fieldsMappingSplits[key] = strings.ToLower(fieldsMappingSplits[key])
			fieldsMappingSplits[key] = strings.Replace(fieldsMappingSplits[key], " ", "", -1)

			//split data into standard fields and client given fields
			fieldSplits := strings.Split(fieldsMappingSplits[key], "=")
			fmt.Println(fieldSplits)
			fmt.Println(len(fieldSplits))

			//unless data is empty or mapping is provided by the client
			if len(fieldSplits) > 1 && fieldSplits[1] != "0" && fieldSplits[1] != "n/a" && fieldSplits[1] != "na" {

				// Remove all the regex from strings
				fieldSplits[0] = reg.ReplaceAllString(fieldSplits[0], "")
				fieldSplits[1] = reg.ReplaceAllString(fieldSplits[1], "")

				//add underscore to primary field if any field is unstructured
				atIndex := strings.Index(fieldSplits[0], "unstructured")
				if atIndex != -1 {
					fieldSplits[0] = fieldSplits[0][:atIndex] + "_" + fieldSplits[0][atIndex:]
					fieldSplits[1] += "_unstructured"
					dataType = "text"
				}
	*/
	/*
		//check to add .keyword to the searchable fields
		//not needed if the mapping is already implemented the right way
		if dataType == "keyword" {
			fieldSplits[0] += ".keyword"
		}*/

	/*			//create the final map
			var tempStringSlice []string
			tempStringSlice = append(tempStringSlice, fieldSplits[1])
			tempStringSlice = append(tempStringSlice, dataType)
			finalMapping[fieldSplits[0]] = tempStringSlice
			//finalMapping[fieldSplits[0]] = append(finalMapping[fieldSplits[0]], fieldSplits[1], dataType)
		}

		//breakdown := []byte(val)
	}
	fmt.Println(finalMapping)

	jsonFileVis, err := os.Open(filepath.Join(currentPath, "savedObjects-visualization.json"))
	if err != nil {
		fmt.Println(err)
	}
	jsonFileIndP, err := os.Open(filepath.Join(currentPath, "savedObjects-indexPattern.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFileVis.Close()
	defer jsonFileIndP.Close()

	// read our opened xmlFile as a byte array.
	byteValueVis, _ := ioutil.ReadAll(jsonFileVis)
	byteValueIndP, _ := ioutil.ReadAll(jsonFileIndP)

	// Declared an empty interface of type Array
	var resultVis []map[string]interface{}
	var resultIndP []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValueVis), &resultVis)
	json.Unmarshal([]byte(byteValueIndP), &resultIndP)

	//change index Values in vis savedObject
	for keyVis, refVis := range resultVis {
		attrVis := refVis["attributes"].(map[string]interface{})
		//str := attrVis["title"].(string)
		visStateVis := attrVis["visState"].(string)

		//visStateVis = strings.Replace(visStateVis, "\\", "", -1)
		//fmt.Println("\\\"field\\\":\\\"")
		indexBeginning := 0
		fmt.Println(strings.Count(visStateVis, "field"))
		for allFields := strings.Count(visStateVis, "field"); allFields > 0; allFields-- {
			if strings.Contains(visStateVis, "\"field\":\"") {

				//indexname
				start := "\"field\":\""
				end := "\""
				s := strings.Index(visStateVis[indexBeginning:], start)
				if s == -1 {
					break
				}
				s += len(start) + indexBeginning
				e := strings.Index(visStateVis[s:], end)
				if e == -1 {
					break
				}
				indexBeginning = s
	*/ /*println("At index:", s)
	println("char is:", string(visStateVis[s]))
	println("chars are:", visStateVis[s:])
	println("At index:", e)
	println("char is:", string(visStateVis[s+e-1]))
	fmt.Println(s, e)*/
	/*				fieldName := visStateVis[s : s+e]
				fmt.Println(fieldName)
				finalString := visStateVis
				for keyMap, valMap := range finalMapping {
					if keyMap == fieldName {
						fmt.Printf("keymap matched fieldname %s and %s. \t", keyMap, fieldName)
						fmt.Println("replaced with", valMap[0])
						finalString = visStateVis[:s] + valMap[0] + visStateVis[s+e:]
					}
				}
				visStateVis = finalString
				attrVis["visState"] = finalString
				refVis["attributes"] = attrVis
				resultVis[keyVis] = refVis
				//midString := visStateVis[:s] + visStateVis[s+e:]
				//fmt.Println("Final visStateVis is:", visStateVis)

			}
		}

		//time.Sleep(5 * time.Second)
	}

	//
	//change index Values in indexPattern savedObject
	for keyIndP, refIndP := range resultIndP {
		attrIndP := refIndP["attributes"].(map[string]interface{})
		//str := attrIndP["title"].(string)
		fieldsIndP := attrIndP["fields"].(string)

		indexBeginning := 0
		fmt.Println(strings.Count(fieldsIndP, "\"name\""))
		for allFields := strings.Count(fieldsIndP, "\"name\""); allFields > 0; allFields-- {
			if strings.Contains(fieldsIndP, "\"name\":\"") {

				//indexname
				start := "\"name\":\""
				end := "\""
				s := strings.Index(fieldsIndP[indexBeginning:], start)
				if s == -1 {
					break
				}
				s += len(start) + indexBeginning
				e := strings.Index(fieldsIndP[s:], end)
				if e == -1 {
					break
				}
				indexBeginning = s
	*/ /*println("At index:", s)
	println("char is:", string(fieldsIndP[s]))
	println("chars are:", fieldsIndP[s:])
	println("At index:", e)
	println("char is:", string(fieldsIndP[s+e-1]))
	fmt.Println(s, e)*/
	/*				fieldName := fieldsIndP[s : s+e]
				fmt.Println(fieldName)
				finalString := fieldsIndP
				for keyMap, valMap := range finalMapping {
					if keyMap == fieldName {
						fmt.Printf("keymap matched fieldname %s and %s. \t", keyMap, fieldName)
						fmt.Println("replaced with", valMap[0])
						finalString = fieldsIndP[:s] + valMap[0] + fieldsIndP[s+e:]
					}
				}
				fieldsIndP = finalString
				attrIndP["fields"] = finalString
				refIndP["attributes"] = attrIndP
				resultIndP[keyIndP] = refIndP
				//midString := fieldsIndP[:s] + fieldsIndP[s+e:]
				//fmt.Println("Final fieldsIndP is:", fieldsIndP)

			}
		}
	}

	byteValueVis, err = json.Marshal(&resultVis)
	if err != nil {
		fmt.Printf("Issue with the visualization Marshal")
		panic(err)
	}
	byteValueIndP, err = json.Marshal(&resultIndP)
	if err != nil {
		fmt.Printf("Issue with the Index Pattern Marshal")
		panic(err)
	}

	ioutil.WriteFile(filepath.Join(currentPath, "savedObjects-visualization2.json"), byteValueVis, 0644)
	if err != nil {
		fmt.Printf("Issue with the visualization json creation")
		panic(err)
	}
	ioutil.WriteFile(filepath.Join(currentPath, "savedObjects-indexPattern2.json"), byteValueIndP, 0644)
	if err != nil {
		fmt.Printf("Issue with the indexPattern json creation")
		panic(err)
	}

	mappingBytes, err := ioutil.ReadFile(filepath.Join(currentPath, "indexMap-sr.json"))
	if err != nil {
		panic(err)
	}

	// Declared an empty interface of type Array
	var resultMapping map[string]interface{}

	fmt.Println("\n\n\n")
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(mappingBytes), &resultMapping)
	//fmt.Println("resultMapping[\"mappings\"]:", resultMapping["mappings"])
	mapMapping := resultMapping["mappings"].(map[string]interface{})
	propMapping := mapMapping["properties"].(map[string]interface{})
	for propKey := range propMapping {
		//fmt.Println("propKey:", propKey)
		for keyMap, valMap := range finalMapping {
			if keyMap == propKey {
				fmt.Printf("keymap matched fieldname %s and %s. \t", keyMap, propKey)
				fmt.Println("replaced with", valMap[0])
				tempMap := make(map[string]string)
				fmt.Println(valMap[1])
				tempMap["type"] = valMap[1]
				//propMapping[propKey] = tempMap
				delete(propMapping, propKey)
				propMapping[valMap[0]] = tempMap
				//mapMapping["properties"] =
			}
		}

	}
	mapMapping["properties"] = propMapping
	resultMapping["mappings"] = mapMapping

	byteValueMapping, err := json.Marshal(&resultMapping)
	if err != nil {
		fmt.Printf("Issue with the Mapping Marshal")
		panic(err)
	}

	ioutil.WriteFile(filepath.Join(currentPath, "indexMap-sr2.json"), byteValueMapping, 0644)
	if err != nil {
		fmt.Printf("Issue with the indexMapping json creation")
		panic(err)
	}
	//str := attrIndP["title"].(string)

	// Output: [a b c]
	//strings.ToLower(str1)
	//strings.Replace(randomString, " ", "", -1)
	*/ //fmt.Println(row)

	/*h, _ := os.Getwd()
	currentPath := filepath.Join(h, "datafiles-temporary")
	secureHead, _ := strconv.ParseBool((os.Getenv("secureHead"))) //6
	eUser := (os.Getenv("eUser"))                                 //6
	ePassword := (os.Getenv("ePassword"))                         //6
	elasticClusterIP := (os.Getenv("elasticClusterIP"))           //6
	dataSize, _ := strconv.Atoi((os.Getenv("dataSize")))          //6
	//elasticPass := (os.Getenv("elasticPass"))            //6

	var gbSize = float64(0)
	//time.Sleep(5 * time.Minute)
	head := "http://"
	if secureHead == true {
		head = "https://"
	}

	/////////sync sequence for dataPush

	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(2)

	//go dataSetResIncident(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	//go dataSetResEvents(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	//go dataSetResKpisHistory(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	go dataSetIncident(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	go dataSetProblem(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	//go dataSetHealth(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	//go dataSetChange(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)
	//go dataSetAsset(&wg, gbSize, dataSize, currentPath, head, eUser, ePassword, elasticClusterIP)

	wg.Wait()
	*/
}

func dataSetResEvents(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData8resevt = [][]string{
		{"alertgroup", "alertkey", "bacid", "component", "componenttype", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "datacenter", "eventtype", "firstoccurrence", "health_status", "hostname", "lastmodified", "node", "provider_account", "servername", "serverserial", "severity", "source_type", "subcomponent", "summary", "tenant_id", "ticketnumber"},
	}

	//sample strings resEvt
	var resEvtStartString = []string{"alertgroup", "alertkey", "bacid", "component", "componenttype", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "datacenter", "eventtype", "firstoccurrence", "health_status", "hostname", "lastmodified", "node", "provider_account", "servername", "serverserial", "severity", "source_type", "subcomponent", "summary", "tenant_id", "ticketnumber"}
	var resAlertGroup = []string{"CvatFgnghf", "VGZ_AG_Ybtvpny_Qvfx", "VGZ_XIZ_FREIRE_QNGNFGBER"}
	var resEvtAlertKey = []string{"apbcvatcebor:PYBHQ", "npe_fieqfse_tizj_rfk", "npe_qfc_8agp_fgq"}
	var resEvtComponent = []string{"VMwareESX", "Windows", "NodeAvail"}
	var resEvtHealthStatus = []string{"Healthy", "Critical", "Warning"}
	var resEvtServerName = []string{"ACME CORPD0P0ALLP", "DTID0P0ALLP"}

	//var resEvtAssignGroup = []string{"HAXABJA", "NOPE-P-ENGRFTNE", "NOPE-P-JVMEFTNE", "NOPE-P-JVMSBTNE", "NOPE-P-YQOJFHCG", "NOPE-V-FDY", "NOPE-V-FGEOBY", "NOPE-V-FLONFR", "NOPE-V-FUNERCG", "NOPE-V-GJFCEBQHPGFHCC", "NOPE-V-GNZFLF", "NOPE-V-HFIVEHF", "NOPE-V-IZJNER", "NOPE-V-JVERYRFF", "NOPE-V-JZGVIBYV", "NOPE-V-MYVAHKFLF", "NOPE-V-NHGBOBY", "NOPE-V-PFBCFVAQ", "NOPE-V-QO7FLOBY", "NOPE-V-QO7QOOBY", "NOPE-V-RKPUNATR", "NOPE-V-VABENPYOBY", "NOPE-V-VAGFZ", "NOPE-V-VAQFBYNEVF", "NOPE-V-VAQJVAGRY", "NOPE-V-VAQNVK", "NOPE-V-VAQYVAHK", "NOPE-V-VOZPYQ-AJ.SVERJNYY", "NOPE-V-VOZPYQ-FN.JVAQBJF", "NOPE-V-VOZPYQ-FN.YVAHK", "NOPE-V-VOZPYQ-IZJNER", "NOPE-V-VOZPYQ-ONPXHC", "NOPE-V-VZFLF", "NOPE-V-ZDOBY", "NOPE-V-ZIFBCFVA", "NOPE-V-ZQNCFCG"}
	//var resEvtAutoGen = []string{"Y", "N"}
	var resEvtContextApp = []string{"\"ACME3_DC_Application1\",\"ACME3_DC_Application27\",\"ACME3_DC_Application38\"", "\"ACME3_DC_Application10\",\"ACME3_DC_Application30\",\"ACME3_DC_Application41\"", "\"ACME3_DC_Application11\",\"ACME3_DC_Application30\",\"ACME3_DC_Application42\"", "\"ACME3_DC_Application14\",\"ACME3_DC_Application30\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application15\",\"ACME3_DC_Application30\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application16\",\"ACME3_DC_Application31\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application17\",\"ACME3_DC_Application31\",\"ACME3_DC_Application44\"", "\"ACME3_DC_Application18\",\"ACME3_DC_Application31\",\"ACME3_DC_Application44\"", "\"ACME3_DC_Application19\",\"ACME3_DC_Application31\",\"ACME3_DC_Application45\"", "\"ACME3_DC_Application2\",\"ACME3_DC_Application22\",\"ACME3_DC_Application27\",\"ACME3_DC_Application39\"", "\"ACME3_DC_Application2\",\"ACME3_DC_Application27\",\"ACME3_DC_Application39\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application31\",\"ACME3_DC_Application46\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application32\",\"ACME3_DC_Application34\",\"ACME3_DC_Application47\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application32\",\"ACME3_DC_Application46\"", "\"ACME3_DC_Application21\",\"ACME3_DC_Application34\",\"ACME3_DC_Application48\"", "\"ACME3_DC_Application21\"", "\"ACME3_DC_Application22\"", "\"ACME3_DC_Application23\"", "\"ACME3_DC_Application24\"", "\"ACME3_DC_Application25\"", "\"ACME3_DC_Application26\",\"ACME3_DC_Application28\",\"ACME3_DC_Application40\",\"ACME3_DC_Application5\"", "\"ACME3_DC_Application26\"", "\"ACME3_DC_Application27\",\"ACME3_DC_Application3\",\"ACME3_DC_Application40\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application4\",\"ACME3_DC_Application40\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application40\",\"ACME3_DC_Application6\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application41\",\"ACME3_DC_Application7\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application41\",\"ACME3_DC_Application8\"", "\"ACME3_DC_Application29\",\"ACME3_DC_Application41\",\"ACME3_DC_Application9\"", "\"ACME3_DC_Application34\"", "\"ACME3_DC_Application35\"", "\"ACME3_DC_Application36\"", "\"ACME3_DC_Application37\"", "\"ACME3_DC_Application38\"", ""}
	var resEvtContextEnv = []string{"UNKNOWN", "DEVELOPMENT", "PRE_PRODUCTION", "DEVELOPMENT", "PRODUCTION", "DEVELOPMENT", "STAGING", "UNKNOWN", "TEST"}
	var resEvtTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var resEvtHostName = []string{"afkznantre", "awccrrfk56", "awccrrfk57", "awccrrfk58", "awccrrfk59", "awcfvasoxcgfz56", "awicegvopyk56", "awicjpherncc56", "cc-ermcnex", "g6cfprapoffdy56", "inosf58", "ipf7", "ipragre-jqp59", "ipragre-qny65", "j59icefueegz56", "j59icefueegz57", "j59icefueica56", "j59icefuencz56", "j59icefueose56", "j59icefueose57", "j59icefueqaf56", "j59icefueqaf57", "j59icefueqbc56", "j59icefueqbc57", "j59icefueryx51", "j59icefueryx56", "j59icefueryxzba56", "j59icefuexo56", "j59icefuexo57", "j59icefueybt56", "j59icefueybt57", "j59icefuezckl56", "j59icefuezy56", "j59iceoxcgfz56", "j59icerpzjro56", "j59icerpzjro57", "j59icerpzncc50", "j59icerpzncc56", "j59icerpzncc57", "j59icerpzncc58", "j59icerpzncc59", "j59icerpzncc65", "j59icerpznrzc50", "j59icerpznrzc51", "j59icerpznrzc56", "j59icerpznrzc57", "j59icerpznrzc58", "j59icerpznrzc59", "j59icerpznrzn56", "j59icerpznrzn57", "j59icerpzpnu56", "j59icerpzpnu57", "j59icerpzqvep56", "j59icerpzqvep57", "j59icerpzrqo56", "j59icerpzrqo57", "j59icerpzrqo58", "j59icerpzrqo59", "j59icjfrpwzc56", "j59icjfuefcir56", "j59icjfuenq56", "j59icjfuenq57", "j59icjfuewzc56", "j59icjrpzngg56", "j59icjrpzwzc56", "j59igefuerqo56", "j59igerpzjro56", "j59igerpzjro57", "j59igerpzncc52", "j59igerpzncc56", "j59igerpzncc57", "j59igerpzncc58", "j59igerpzncc59", "j59igerpznrz50", "j59igerpznrz51", "j59igerpznrz52", "j59igerpznrz53", "j59igerpznrz54", "j59igerpznrz56", "j59igerpznrz57", "j59igerpznrz58", "j59igerpznrz68", "j59igerpzrqo56", "j59igerpzrqo58", "j59igerpzrqo59", "j59igjrpzwzc56", "jud-jvmpbz-ceq7", "jud-jvmpbz-oqp6", "nmcfnipbecqp56", "nmcfnipbecqp57", "o6cfnicebkgnq56", "o6cfnifrcqz56", "o6cfnifrcva56", "o6cfniibygfdy56", "o6cfniieffdy57", "o6cfninccpbe56", "o6cfnisfpncc56", "o6cfnopragjro56", "o6cnierpbancc56", "o6icynopucncc56", "o6ihynopucncc56", "o78mcejvmncc56", "o78mcejvmncc57", "o78mcejvmzba56", "o78mcejvmzba57", "o78mhejvmncc56", "o7cffnozyttgj57", "o7cfnibz0fbn56", "o7cfnibz0fbn57", "o7cfniepapgk57", "o7cfnipfnszna56", "o7cfnipfuejgf56", "o7cfnippngncc56", "o7cfnipzyttgj53", "o7cfnipzyttgj57", "o7cfnohngfdy56", "o7cfnohngjro56", "o7cfnohngncc56", "o7cfnotbz0qo56", "o7cfnotbz0qo57", "o7cfnotbz0qo58", "o7cfnotgjvncc56", "o7cfnotouq56", "o7cfnotzaqncc56", "o7cfnzqzyttgj57", "o7cfongtqftgj68", "o7cfpbeohfncc58", "o7cfpbeuhzfdy56", "o7cfpbeuhzvak56", "o7qfni9fbntom", "o7qfnibz9ecg56", "o7qfnibz9fbn56", "o7qfnibz9fbnm6", "o7qfnibz9fdy56", "o7qfnohngztg56", "o7qfnotbz8qo56", "o7qfnotbz9qo56", "o7qfnotbz9qo57", "o8ccecebben66", "o8ccecebben67", "o8ccecebben68", "o8ccecebncc68", "o8ccecebncc69", "o8ccecnhben56", "o8ccecrzben56", "o8ccegjben56", "o8ccegjben57", "o8ccegjvncc56", "o8ccegjvncc57", "o8ccenqoben56", "o8ccenqoben57", "o8ccepzfben56", "o8ccepzfncc56", "o8cceqclben56", "o8cceqclncc57", "o8cceqclpqz57", "o8cceqfgben56", "o8cceqfgncc56", "o8cceqfoben56", "o8cceqfoncc56", "o8cceqonben56", "o8cceqonben57", "o8cceqonben58", "o8cceqscben56", "o8cceqyaben56", "o8cceqyancc57", "o8cceqyjben56", "o8cceqypben56", "o8ccerzrncc56", "o8ccerzrncc57", "o8ccexnapfr50", "o8ccjcnhfdy56", "o8ccjfnifsopdq6", "o8ccjfnifsoqve7", "o8ccjfnifsosr56", "o8ccjgcpfpp56", "o8ccjjeyfdy50", "o8ccjjeyfdy57", "o8ccjppnencc56", "o8ccjppnencc57", "o8ccjvasrkz50", "o8ccjvasrkz51", "o8ccjvasrkz56", "o8ccjvasrkz57", "o8ccjvasrkz58", "o8ccrbztfs56", "o8ccrbztfs57", "o8ccrbztfu56", "o8ccrbztfu57", "o8ccrbztfu58", "o8ccrecgrfk56", "o8ccrecgrfk57", "o8ccrfsrfk50", "o8ccrfsrfk51", "o8ccrfsrfk52", "o8ccrfsrfk53", "o8ccrfsrfk56", "o8ccrfsrfk57", "o8ccrfsrfk58", "o8ccrfsrfk59", "o8ccrwhzcfu57", "o8ccrzqzrfk56", "o8ccrzqzrfk57", "o8cfjjeyfdy56", "o8cfjjeyfdy57", "o8cfjjeyfdy59", "o8cfnibz0qgpby6", "o8cfnibz0zd57", "o8cfnibz0zqz57", "o8cfnic3nr56", "o8cfnic3nr57", "o8cfnic3pc56", "o8cfnifzfben56", "o8cfnifzfncc56", "o8cfnigecben56", "o8cfniic3jro56", "o8cfniic3jro57", "o8cfnijngjf57", "o8cfnipxnapfr56", "o8cfnipynben56", "o8cfnjmyqncc57", "o8cfnjmyqncc58", "o8cfnjvmjf57", "o8cfnjvmjf58", "o8cfnjvmjf59", "o8cfnotfgtncc56", "o8cfnotfgtqo56", "o8cfnotvas56", "o8checebben66", "o8checebben67", "o8checebben68", "o8checebncc60", "o8checebncc66", "o8checebncc67", "o8checebncc68", "o8checebncc69", "o8checnhben56", "o8checrzben56", "o8cheenhncc56", "o8chenebflo68", "o8chenebncc66", "o8chenebncc68", "o8chenebncc69", "o8chenqoben56", "o8chenqoben57", "o8chepzfben56", "o8chepzfncc56", "o8cheqanpqz51", "o8cheqanpqz53", "o8cheqanpqz58", "o8cheqclben56", "o8cheqclncc57", "o8cheqclpqz56", "o8cheqfgben56", "o8cheqfoben56", "o8cheqfoncc57", "o8cheqonben56", "o8cheqonben57", "o8cheqscazd56", "o8cheqscben56", "o8cheqyaben56", "o8cheqyjben56", "o8cheqyjncc56", "o8cheqyjncc57", "o8cheqypben56", "o8cheqypncc56", "o8cheqypncc57", "o8cherzrncc56", "o8chjgnfncc56", "o8chjjeyfdy56", "o8chjjeyfdy57", "o8chjqclfdy56", "o8chrbztfs56", "o8chrbztfs57", "o8chrzqzrfk56", "o8cjvasfxyz58", "o8cjvasfxyz59", "o8cjylpcby50", "o8cjylpcby59", "o8cjylppzf50", "o8cjylppzf59", "o8cqenebflo66", "o8cqepynben56", "o8cqeqscazd57", "o8ffprqjc-qo6", "o8ffprqjc-qo7", "o8hfnic3nr56", "o8hfnic3nr57", "o8hfnic3pc56", "o8hfnic3pc57", "o8hfnic3pyqz56", "o8hfnic3pyqz57", "o8hfnifzfben56", "o8hfnifzfncc56", "o8hfniic3fdy56", "o8hfniic3jro56", "o8hfnipzfben56", "o8hfnipzffdy56", "o8icebztfvzncc56", "o8icebztncc50", "o8icebztncc51", "o8icebztncc56", "o8icebztncc57", "o8icebztncc58", "o8icebztncc59", "o8icepngqon56", "o8icepngqon57", "o8icjavegnc56", "o8icjavprnc56", "o8icjavprqo56", "o8icjavprratr56", "o8icjbzincc58", "o8icjcnlyffdy56", "o8icjecgjfncc56", "o8icjecgjfncc57", "o8icjecgjfncc58", "o8icjecgjfncc59", "o8icjecgpzncc56", "o8icjecgpzncc57", "o8icjgrz56", "o8icjjeyjro54", "o8icjnepuefei56", "o8icjnepuejro56", "o8icjnoorffdy57", "o8icjnotbbqtc53", "o8icjnotbbqtp52", "o8icjnotcegt56", "o8icjnotpgkkn50", "o8icjnotpgkkn51", "o8icjpgkkra50", "o8icjpgkkra51", "o8icjpgkkra52", "o8icjpgkkra54", "o8icjpgkkra56", "o8icjpgkqo56", "o8icjpnvff56", "o8icjpoffdy56", "o8icjqnvj57", "o8icjsbqscy56", "o8icjvasacf56", "o8icjvasirnz56", "o8icjvasjro56a", "o8icjvasjro57", "o8icjvasrkqnt56", "o8icjvasrkrqt59", "o8icjwhzcubfg56", "o8icjwhzcubfg57", "o8icyfrpfpna58", "o8ifjjeyjro56", "o8ifjjeyjro57", "o8ihebztncc56", "o8ihebztncc57", "o8ihebztyqnc56", "o8ihezqztoy56", "o8ihjhvcfdy56", "o8ihjhvcjro56", "o8ihjhvcjro57", "o8ihjjeyncc56", "o8ihjnotbbqof50", "o8ihjnotbbqtc57", "o8ihjnotbbqtp56", "o8ihjnotbbqtz58", "o8ihjpbecqp56", "o8ihjylpncc56", "o8ihjylpncc57", "o8iqebztncc56", "o8iqeqonben56", "o8iqjdnncc56", "o8iqjdnqo56", "o8iqjohffdy-g", "o8iqjohffdy56", "o8iqjzyttgj56", "o8qfnibz9hvn56", "o8qfnibz9zqz56", "o8qfniegpben56", "o8qfnigvzncc56", "o8qfnipynben56", "o8qfnipzoncc56", "o8qfnotqrincc56", "o8qfnotvas56", "o8qfnotvas57", "o8qfnotvasyqz56", "o8qjfcecg56", "o8tmcfniperi56", "o8tmcfniperi57", "oc6cfvasrfk56", "oc6cfvasrfk57", "ocvcfvasfwf56", "oemcfninebfym56", "oemcfnipfueqon56", "oemcfnipfuezfp56", "oemcfniplyqqo56", "ofmdfnipfuezfp56", "ofmdfnipfuezfp57", "ofmjepqri57", "ofmqfnipfueqri56", "ofmqfnipfueqri59", "ofmqfniprqjqri56", "ojud-ebire-jjxf", "onopejf050", "onopenc053", "onopenc065", "onopenc066", "onopenc553", "onopenc554", "onopenc561", "onopenc562", "onopenc564", "onopenc565", "onopenc567", "onopenc568", "onopenc569", "onopenc575", "onopenc576", "oqcvasgqztgjl56", "or7qfnisri8frp58", "otmdfnipfuefei56", "otmqfnipfuefei56", "ov6cfvasip56", "ov6cfvasipqo56", "ovcftevaspgk56", "ovcftevaspgk57", "ovcftevaspgk58", "ovcfvasavz96", "ovcfvasegqp57", "ovcfvasfucjro56", "ovcfvasfucncc56", "ovcfvasgfz-yna-serr", "ovcfvasgqz56", "ovcfvasjvaf56", "ovcfvasqqp57", "ovcvasgqztgjl56", "ovqppraqnagqzm7", "ovqppraqnagqzm8", "ovqzmvfn57", "ovvcfqupc59", "pnopenc053", "pnopenc054", "pnopenc060", "pnopenc066", "pnopenc068", "pnopenc069", "pnopenc553", "pnopenc554", "pnopenc561", "pnopenc564", "pnopenc575", "pnopeqo050", "pnopeqo059", "pnopeqo550", "pnopeqo551", "pnopeqo559", "q65icefueegz56", "q65icefueegz57", "q65icefueica56", "q65icefuencz56", "q65icefuencz57", "q65icefueose56", "q65icefueose57", "q65icefueqaf56", "q65icefueqaf57", "q65icefueqbc56", "q65icefueybt56", "q65icefueybt57", "q65icefuezckl56", "q65iceoxcgfz56", "q65icerpzjns57", "q65icerpzjro56", "q65icerpzjro57", "q65icerpzncc50", "q65icerpzncc51", "q65icerpzncc52", "q65icerpzncc53", "q65icerpzncc54", "q65icerpzncc56", "q65icerpzncc57", "q65icerpzncc58", "q65icerpzncc59", "q65icerpzncc65", "q65icerpznczp56", "q65icerpznrzc50", "q65icerpznrzc51", "q65icerpznrzc56", "q65icerpznrzc57", "q65icerpznrzc58", "q65icerpznrzc59", "q65icerpzpnu56", "q65icerpzpnu57", "q65icerpzqvep56", "q65icerpzqvep57", "q65icerpzrqo50", "q65icerpzrqo56", "q65icerpzrqo57", "q65icerpzrqo58", "q65icerpzrqo59", "q65icjfrpwzc56", "q65icjfuefcir56", "q65icjfuenq56", "q65icjfuenq57", "q65icjfueqnvc56", "q65icjfueqnvc57", "q65icjfuewzc56", "q65icjrpzwzc56", "q65igerpzjns56", "q65igerpzjro56", "q65igerpzjro57", "q65igerpzncc56", "q65igerpzncc57", "q65igerpzncc58", "q65igerpzncc59", "q65igerpznczp56", "q65igerpznrz50", "q65igerpznrz51", "q65igerpznrz52", "q65igerpznrz53", "q65igerpznrz54", "q65igerpznrz56", "q65igerpznrz57", "q65igerpznrz58", "q65igerpznrz59", "q65igerpznrz65", "q65igerpznrz66", "q65igerpznrz67", "q65igerpzrqo50", "q65igerpzrqo56", "q65igerpzrqo57", "q65igerpzrqo58", "q65igerpzrqo59", "q65igjrpzwzc56", "qt9bfc7e6iw", "r7cfnipfvtpnc56", "r7cfnipfvtpnc57", "r7cfnipoymqri59", "r7cfnipsygjgf57", "r7cfnipvasuvf54", "r7cfnipvasuvf66", "r7cfnotbz0ncc56", "r7cfnotbz0ncc57", "r7cfnotbz0vag56", "r7cfnotbz0vag57", "r7cfpbecqzmqp57", "r7hfnipfvtpnc56", "r7qfnipfvtpnc56", "r7qfnisri8jf56", "r7qfnotbz8ncc56", "r7qfnotbz9ncc56", "r7qfnotbz9ncc57", "r7qfnotbz9vag56", "r7qfnotbz9vag57", "r7qfnotyzptgj56", "rfkv0", "rfkv1", "rfkv2", "rfkv5", "rfkv6", "rfkv7", "rfkv8", "rfkv9", "ugqp56-221qp154", "ugqp57-221qp154", "zif6", "zif7"}
	//var resEvtSituation = []string{"apbcvatcebor", "apbcvatcebor:PYBHQ", "cvat", "FE_FAZC BGZN sybbq vf qrgrpgrq:Fcyhax Ebire Nyreg - BGZN sybbq:NOPE-V-PFBCFVAQ", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: gfg-cevprznantre.noenp.arg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: jvmneqthv-hng7.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: NOPEVagreanyVffhvatPN  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: NOTFFY  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: CEBF-CEVPVAT  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: EBIRE-ARKGTRA  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: gfg-cevprznantre-jrfgrea  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-anz-qsc-abegurnfg.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-anz-qsc-jrfg.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-zbovyrnccf.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: jvmneqthv-cnlyrffpne  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: jvmneqthvohqtrg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: OCZHNG.NIVFOHQTRG.PBZ  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: qrsnhyg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: vevf-hng.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=1:freireAnzr=JQP-DN-644-NOTPBZD (65.679.1.644: 64266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=4:freireAnzr=JQP-DN-644-Ohqtrg (65.679.1.644: 63266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=6:freireAnzr=Cbfgterf Ragrecevfr Znantre Freire (672.5.5.6: 62286):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=65:freireAnzr=NOTPBZ-DN-Qnyynf-643 (65.679.25.643: 64266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=66:freireAnzr=NIVF-DN-Qnyynf-644 (65.679.25.644: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=11.3:inyhr=30.3:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Znk pbaarpg", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=4:freireAnzr=Ohqtrg-DN-Qnyynf-643 (65.679.25.643: 63266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=18.8888888888888:inyhr=32.0:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngv", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=4:freireAnzr=Ohqtrg-DN-Qnyynf-643 (65.679.25.643: 63266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=29.0:inyhr=35.0:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Znk pbaarpgv", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=18.3888888888888:inyhr=34:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Znk p", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=25.0:inyhr=37:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Znk pbaarpgvba = 155", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=1:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = ragrecevf\"", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=3:freireAnzr=Qnyynf-Cebq-643-NOTPBZ (65.679.12.643: 4266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRG nccyvpngvba_anzr = 'RagrecevfrQ", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=4:freireAnzr=Qnyynf-Cebq-643-Ohqtrg (65.679.12.643: 3266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=1:inyhr=6:cerivbhfFgnghf=YBJ:fgnghf=PYRNE:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = ra\"", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=6:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = rag\"", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=83:inyhr=5:cerivbhfFgnghf=UVTU:fgnghf=PYRNE:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=1:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG bvq, sbezng_glcr(bvq, AHYY)", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr-va-genafnpgvba fgngr:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=5:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRG FRFFVBA frnepu_cngu", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr-va-genafnpgvba fgngr:freireVQ=65:freireAnzr=NOTPBZ-DN-Qnyynf-643 (65.679.25.643: 64266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=1:inyhr=5:cerivbhfFgnghf=YBJ:fgnghf=PYRNE:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=JQP-DN-643-Ntrag:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=1.9561405555555555:inyhr=45.6020617555555555:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=JQP-DN-643-Ntrag:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=69.1812767555555555:inyhr=25.3029871555555555:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=73.9669835555555555:inyhr=34.7356892555555555:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=4896:inyhr=2447:cerivbhfFgnghf=UVTU:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=9870:inyhr=4611:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr \"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=8753:inyhr=1986:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr fv\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=4:freireAnzr=Qnyynf-Cebq-643-Ohqtrg (65.679.12.643: 3266):guerfubyqInyhr={261,364,476}:cerivbhfInyhr=6876:inyhr=6876:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = ohqrpbzc", "Qngnonfr \"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=65:freireAnzr=Qnyynf-Cebq-643-Nivf (65.679.12.643: 2266):guerfubyqInyhr={6579,6773,6081}:cerivbhfInyhr=6437:inyhr=6437:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzc", "Qngnonf\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=66:freireAnzr=NIVF-DN-Qnyynf-644 (65.679.25.644: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=8565:inyhr=1355:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnon\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={6579,6773,6081}:cerivbhfInyhr=8717:inyhr=8717:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzc", "Qngnonf\"", "FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={954,067,169}:cerivbhfInyhr=993:inyhr=993:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=2087:inyhr=4299:cerivbhfFgnghf=ZRQVHZ:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr fvm\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=04.2193040134547:inyhr=47.9330140178220:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-643", "Svyr flfgrz = /\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=39.3635304820:inyhr=46.561873670:cerivbhfFgnghf=YBJ:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-643", "Svyr flfgrz = 616.71.43.\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=8:ntragAnzr=rcnf66-dn-644:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=23.4100304820:inyhr=46.561873670:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-644", "Svyr flfgrz = 616.71.4\"", "FIP:EGRZFgnghf:gvizba:VGZ:7", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):56:54:52.712391-75755651/PN70F", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):66:72:74.739910-75646785/PN70F", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):76:51:64.157393-75646786/PN70F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(275):75755676-54:72:85.450360/PN50F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(VRN935R):75755671-58:55:73.856007/PN50F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(VRN935R):75755671-58:59:51.436116/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):55:51:90.539761-75755671/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):62:53:59.938049-75755668/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):63:67:63.398179-75755669/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):65:75:60.979634-75755668/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):69:73:68.992363-75755667/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):75:78:54.628604-75755669/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):76:86:09.668552-75755658/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):76:95:95.840055-75755658/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755686-55:82:96.515217/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755686-55:83:53.319918/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755756-78:84:99.304337/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755757-55:56:02.505872/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755758-66:02:92.466402/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755758-67:58:80.154069/PN70F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(008):75755660-51:70:95.139031/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(018):75755660-51:89:71.183703/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(019):75755660-51:81:85.071212/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(059):75755660-51:55:85.394300/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(061):75755660-51:66:00.917919/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(092):75755660-51:77:53.888751/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP01722) ZftVq(QFAW666R):75755652-52:80:60.092573/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP01722) ZftVq(QFAW666R):75755664-51:57:83.593827/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP50606) ZftVq(QFAW666R):75646761-51:57:09.440147/PN50F", "NGZ_:WboAnzr(OHQQYVG) Ahzore(FGP63730) ZftVq(QSF5399V):75755679-50:90:55.328696/PN50F", "NGZ_:WboAnzr(OHQQYVG) Ahzore(FGP68216) ZftVq(QSF5399V):75755673-54:97:81.695926/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAG055V):75755676-59:62:88.303018/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAW666R):75755661-60:95:97.533087/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAW666R):75755675-57:02:92.126541/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP67686) ZftVq(QFAG055V):75755759-65:91:06.262786/PN50F", "NGZ_:WboAnzr(QFOEZFGE) Ahzore(FGP67673) ZftVq(QFAW666R):75755685-61:59:77.634516/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAG055V):75755676-58:79:50.096956/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAW666R):75755669-66:77:82.870982/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAW666R):75755677-78:95:05.215797/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP50608) ZftVq(QFAW666R):75646760-62:54:63.628294/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP67685) ZftVq(QFAW666R):75755674-52:67:88.671237/PN50F", "NGZ_:WboAnzr(QSUFZNI7) Ahzore(FGP01884) ZftVq(NEP5558V):75646785-66:73:69.855100/PN50F", "NGZ_:WboAnzr(UJFZREO7) Ahzore() ZftVq(QSF6434R):75755672-66:59:74.039318/PN50F", "NGZ_:WboAnzr(VZF) Ahzore(FGP51864) ZftVq(QSF7582):55:66:68.892755-75755671/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755672-54:60:58.527614/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755685-69:55:59.537925/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755758-54:60:58.529562/PN70F", "NGZ_:WboAnzr(VZFNCCY4) Ahzore() ZftVq(6):75755759-54:85:59.659029/PN70F", "NGZ_:WboAnzr(VZFNPCO8) Ahzore() ZftVq(QSF6434R):69:57:83.190071-75755658/PN70F", "npe_abdzte_tzdp_fgq", "npe_ajcbeg_j51s_jva", "npe_bssyvar_ezdp_gvi", "npe_bssyvar_tzfj_fgq", "npe_bssyvar_tzfj_pybhq", "npe_cepoyx_tblj_fgq", "npe_ceppch_8agj_fgq", "npe_ceppch_8agj_fgq_Pyq", "npe_ceppchuv_eymp_gvi", "npe_cepzvf_eymj_gvi_wnmmfz", "npe_cepzvf_eyms_gvi_xghpzn", "npe_cepzvf_khkj_fhabar", "npe_cepzvf_khkj_nrz", "npe_cntsvyr_8agj_fgq", "npe_dqrcgu_dzdp_r7cfnotbz0vag6", "npe_dqrcgu_dzds_o7cfnibz0fbn", "npe_dqrcgu_dzds_r7cfnotbz0ncc", "npe_dqrcgu_dzds_r7cfnotbz0vag", "npe_dqrcgu_dzdz_o7cfnibz0fbn56", "npe_dqrcgu_dzdz_r7cfnotbz0vag57", "npe_dqrcgu_dzdz_r7cfnotbz0vag6", "npe_dquv_ezdj_tqfg_erf56ontohqe", "npe_dquv_mzdp_cebq_flfgrz", "npe_dshyy_dzdj_o8icebztncc", "npe_dshyy_dzdp_o8icebztncc", "npe_dzteceo_mzds_cebq", "npe_efgvzybj_egbp_rpc_erf56c05", "npe_efgvzybj_egbp_rpc_erf56cv7", "npe_efgvzybj_egbp_rpc_erf56cv8", "npe_efgvzybj_egbp_rpc_erf56cv9", "npe_efgvzybj_egbp_rpc_erf56l05", "npe_efgvzybj_egbp_rpc_erf56lv8", "npe_efgvzybj_egbp_rpc_erf56o05", "npe_efgvzybj_egbp_rpc_erf56ov7", "npe_efgvzybj_egbp_rpc_erf56ov8", "npe_efgvzybj_egbp_rpc_erf56ov9", "npe_efgvzybj_egbp_tqfc_erf56cn6", "npe_efgvzybj_egbp_tqfc_erf56cn8", "npe_efgvzybj_egbp_tqfc_erf56cnr", "npe_efgvzybj_egbp_tqfc_erf56cnt", "npe_efgvzybj_egbp_tqfc_erf56lnr", "npe_efgvzybj_egbp_tqfc_erf56lnt", "npe_efgvzybj_egbp_tqfc_erf56on8", "npe_efgvzybj_egbp_tqfc_erf56onr", "npe_efgvzybj_egbp_tqfc_erf56ont", "npe_erderfc_tlaj_jf", "npe_feifgng_tlaj_jf", "npe_feifgng_tlap_jf", "npe_ffyntr_t9fp_ffypreg", "npe_fieepie_tizj_rfk", "npe_fiegfze_tizj_rfk", "npe_fiepch_tizj_rfk", "npe_fieqfse_tizj_rfk", "npe_fiezrz_tizj_rfk", "npe_fiezrz_tizp_rfk", "npe_fip_8agp_fgq", "npe_fip_8agp_fgq_Pyq", "npe_fip_8agp_fgvfip", "npe_fjnccep_fhkj_fgq", "npe_fjnccep_khkj_fgq", "npe_fvgrqja_tugj_uggc", "npe_fvgrsy_tugj_uggc", "npe_fvgrsy_tugp_uggc", "npe_gen_eymj_fgq", "npe_gen_eymj_gvi", "npe_gen_fhkj_fgq", "npe_gen_khkj_fgq", "npe_gffgng_temj_ben", "npe_gfhody_8rkj_fgq", "npe_jpegrkc_8d2j_fgq_i7", "npe_mbz_eymp_fgq_Pyq", "npe_mbz_fhkp_fgq", "npe_mbz_khkp_fgq", "npe_pch_eymj_fgq_pyq", "npe_pch_eymp_cebq", "npe_pch_eymp_fgq", "npe_pch_fhkp_fgq", "npe_pch_khkp_fgq", "npe_pheqcg_tzdp_fgq", "npe_pheqcg_tzds_fgq", "npe_pufgng_ezdp_grfg_jvmzbqgfg", "npe_qfc_8agj_fgq7i7", "npe_qfc_8agp_fgq", "npe_qfc_8agp_fgq_Pyq", "npe_qfc_8ags_fgq", "npe_qfc_8ags_fgq_Pyq", "npe_qobssya_8bdp_fgq", "npe_qofgng_tblp_fgq", "npe_qosvyr_temj_ben", "npe_qovanpg_8bdp_fgq", "npe_qovanpg_temj_fgqdn", "npe_qovanpg_temp_fgq", "npe_qovanpg_temp_fgqdn_rouf", "npe_qovapfg_8bdp_fgq", "npe_qoyef_8rkj_fgq", "npe_qsfgng_temp_fgq", "npe_qvfxfc_temj_ben", "npe_qvfxfc_temp_ben", "npe_qvfxfc_temz_bendn", "npe_reebe_euqp_gvi", "npe_reeybt_fhyj_nonphf", "npe_reeybt_fhyj_nonphf7", "npe_sergofc_temj_ben", "npe_sergofc_temj_ben_rouf", "npe_sergofc_temj_bendn", "npe_sergofc_temj_bendn_rouf", "npe_sergofc_temj_o8ffprqjc", "npe_sergofc_temp_ben", "npe_sergofc_temp_o8ffprqjc", "npe_sergofc_temz_bendn", "npe_sergofc_temz_bendn_rouf", "npe_sff_eymj_fgq", "npe_sff_eymj_fgq_cebq", "npe_sff_eymj_fgq_Pyq", "npe_sff_eymj_gvi_cebq", "npe_sff_eymj_zdz", "npe_sff_eymp_fgq", "npe_sff_eymp_fgq_Pyq", "npe_sff_eymp_gvi_cebq", "npe_sff_eymp_ncc_grfg", "npe_sff_eymp_pyqsvyrcrez", "npe_sff_eyms_fgq", "npe_sff_eyms_fgq_Pyq", "npe_sff_eyms_gvi_cebq", "npe_sff_fhkj_fgq", "npe_sff_fhkj_ROUF", "npe_sff_fhkj_ROUF_gfz", "npe_sff_fhkj_ROUF_grfg", "npe_sff_fhkp_fgq", "npe_sff_fhkp_ROUF", "npe_sff_fhks_fgq", "npe_sff_fhks_ROUF", "npe_sff_khkj_fgq", "npe_sff_khkj_gfzfgt", "npe_sff_khkj_o7cfnotouq56", "npe_sff_khkj_o7qfni9fbntom", "npe_sff_khkj_o7qfnibz9fbn56", "npe_sff_khkj_QN", "npe_sff_khkj_zdz", "npe_sff_khkp_fgq", "npe_sff_khkp_o7qfni9fbntom", "npe_sff_khkp_o7qfnibz9fbn56", "npe_sff_khkp_pfg", "npe_sff_khkp_pfg9", "npe_sff_khkp_zdz", "npe_sff_khks_fgq", "npe_sff_khks_o7qfnibz9fbn56", "npe_sff_khks_zdz", "npe_sofcnpr_temj_bendn", "npe_uggcfg_fugj_uggc", "npe_uggcfg_tugj_uggc", "npe_uogg_eymp_gvi", "npe_uozd_eymp_gvi", "npe_vaf_eymp_fgq", "npe_vaf_fhkp_fgq", "npe_vafgnpg_temj_bendn", "npe_vafgnpg_temp_ben", "npe_vafgnpg_temp_bendn_rouf", "npe_wbosnvy_tbdp_zffdy_cebq", "npe_ybtben_temj_bendn7_rouf", "npe_ybtben_temz_bendn9_rouf", "npe_ybtserr_tblp_fgq", "npe_znkpua_ezdp_cebq_jvmncc56", "npe_znkpua_ezdp_cebq_jvmncc57", "npe_znkpua_ezds_cebq_jvmncc56", "npe_znkpua_ezds_cebq_jvmncc57", "NPF Flagurgvp Gerr Nhgbzngvba Rirag_5", "VGZ:75646786-56"}
	//var resEvtStatus = []string{"CANCELLED", "CLOSED", "INPROG", "PENDING", "QUEUED", "REJECTED", "RESOLVCONF", "RESOLVED", "SLAHOLD"}
	//var resEvtTicketClass = []string{"Application Down", "Backup Missed/Failed", "CPU High Issue", "Database Handler", "Disk Storage Issues", "High Memory & Page File Usage", "ITM/Other Agent", "Job Abends", "MQ Handler", "Non-os Windows Disk Full", "Other Automata", "Server Unavailable", "Service in Alert State", "Swap Space Issue", "Table Space Handler", "Unclassified Actionable", "Windows OS Disk Full", "Windows Service Handler", "Zombie Processes"}

	flag := true
	for gbSize < 200 {
		csvData8resevt = nil
		csvData8resevt = append(csvData8resevt, resEvtStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand alertgroup
			tmpSlice = append(tmpSlice, resAlertGroup[rand.Intn(len(resAlertGroup))])

			//rand alertkey
			tmpSlice = append(tmpSlice, resEvtAlertKey[rand.Intn(len(resEvtAlertKey))])

			//rand BACID
			min := 0200000
			max := 9999999
			randClientID := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand component
			tmpSlice = append(tmpSlice, resEvtComponent[rand.Intn(len(resEvtComponent))])

			//rand componenttype
			tmpSlice = append(tmpSlice, "ComputerSystem")

			//rand context.application
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 0
			max = 4
			randNum := rand.Intn(max-min+1) + min
			collectiveString := ""
			for trav := 0; trav <= randNum; trav++ {
				collectiveString = collectiveString + resEvtContextApp[rand.Intn(len(resEvtContextApp))]
				if trav != randNum {
					collectiveString = collectiveString + ","
				}
			}
			tmpSlice = append(tmpSlice, "["+collectiveString+"]")

			//rand context.environemnt
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\""+resEvtContextEnv[rand.Intn(len(resEvtContextEnv))]+"\"]")

			//rand context.manage
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\"IBM\"]")

			//rand context.team
			min = 1
			max = 9
			randTeam := rand.Intn(max-min+1) + min
			//fmt.Println(randTeam)
			tmpSlice = append(tmpSlice, "[\"acme sales demo"+strconv.Itoa(randTeam)+"\"]")

			//rand TENANT_ID part1
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = resEvtTenantID[rand.Intn(len(resEvtTenantID))]
			}

			//rand correlation_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 0000000
			max = 9999999
			seqno := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "IBM-"+string(TenantIDRune)+"-100000000-"+strconv.Itoa(seqno))

			//rand datacenter
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "IBM")

			//rand eventtype
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "Infrastructure Event")

			//rand firstoccurrence
			ranFstOccur, ranTimeStampFstOccur, _ := randate()
			tmpSlice = append(tmpSlice, ranFstOccur)

			//rand health_status
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resEvtHealthStatus[rand.Intn(len(resEvtHealthStatus))])

			//rand hostname
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resEvtHostName[rand.Intn(len(resEvtHostName))])

			//rand lastmodified
			min = 24
			max = 512
			randRef := rand.Intn(max-min+1) + min
			timeStampRef := ranTimeStampFstOccur.Add(time.Duration(randRef) * time.Minute)
			timeStampA := timeStampRef.Format("Jan 02, 2006")
			timeStampB := timeStampRef.Format("15:04:05")
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")

			//rand node
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand provider_account
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand servername
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resEvtServerName[rand.Intn(len(resEvtServerName))])

			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////
			//////////
			//////////////////
			//////////////////////////

			//create the final slice
			csvData8resevt = append(csvData8resevt, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData8resevt)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData8resevt)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/resource_incidents_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData8resevt)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "resource_incidents_temp.csv") + " > " + filepath.Join(currentPath, "resource_incidents_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted asset")
		}

		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData8resevt)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "resource_incidents_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length for resEvt", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "resource_incidents", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-resource_incidents.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-resource_incidents.json", head, eUser, ePassword, elasticClusterIP, "resource_incidents")

			//cleanup the file
			cleanup(currentPath, "finalOutput-resource_incidents.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-resource_incidents.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-resource_incidents.json", head, eUser, ePassword, elasticClusterIP, "resource_incidents")

		//cleanup the file
		cleanup(currentPath, "finalOutput-resource_incidents.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/resource_incidents/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			//fmt.Println("Response for increasing limit", response)

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/resource_incidents/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("resource_incidents index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		//time.Sleep(5 * time.Second)

		////
		/////
		//////
	}

	defer wg.Done()
}

func dataSetResIncident(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData7resinc = [][]string{
		{"actionable", "assignment_group", "autogenerated", "bacid", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "created", "data_source", "datacenter", "description", "hostname", "last_refresh_dttm", "modify_dttm", "number", "priority", "provider_account", "resolution", "resolved", "situation", "source_type", "status", "symptom", "tenant_id", "ticket_class", "ticket_classification"},
	}

	//sample strings resInc
	var resIncStartString = []string{"actionable", "assignment_group", "autogenerated", "bacid", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "created", "data_source", "datacenter", "description", "hostname", "last_refresh_dttm", "modify_dttm", "number", "priority", "provider_account", "resolution", "resolved", "situation", "source_type", "status", "symptom", "tenant_id", "ticket_class", "ticket_classification"}
	var resIncActionable = []string{"Y", "N"}
	var resIncAssignGroup = []string{"HAXABJA", "NOPE-P-ENGRFTNE", "NOPE-P-JVMEFTNE", "NOPE-P-JVMSBTNE", "NOPE-P-YQOJFHCG", "NOPE-V-FDY", "NOPE-V-FGEOBY", "NOPE-V-FLONFR", "NOPE-V-FUNERCG", "NOPE-V-GJFCEBQHPGFHCC", "NOPE-V-GNZFLF", "NOPE-V-HFIVEHF", "NOPE-V-IZJNER", "NOPE-V-JVERYRFF", "NOPE-V-JZGVIBYV", "NOPE-V-MYVAHKFLF", "NOPE-V-NHGBOBY", "NOPE-V-PFBCFVAQ", "NOPE-V-QO7FLOBY", "NOPE-V-QO7QOOBY", "NOPE-V-RKPUNATR", "NOPE-V-VABENPYOBY", "NOPE-V-VAGFZ", "NOPE-V-VAQFBYNEVF", "NOPE-V-VAQJVAGRY", "NOPE-V-VAQNVK", "NOPE-V-VAQYVAHK", "NOPE-V-VOZPYQ-AJ.SVERJNYY", "NOPE-V-VOZPYQ-FN.JVAQBJF", "NOPE-V-VOZPYQ-FN.YVAHK", "NOPE-V-VOZPYQ-IZJNER", "NOPE-V-VOZPYQ-ONPXHC", "NOPE-V-VZFLF", "NOPE-V-ZDOBY", "NOPE-V-ZIFBCFVA", "NOPE-V-ZQNCFCG"}
	var resIncAutoGen = []string{"Y", "N"}
	var resIncContextApp = []string{"\"ACME3_DC_Application1\",\"ACME3_DC_Application27\",\"ACME3_DC_Application38\"", "\"ACME3_DC_Application10\",\"ACME3_DC_Application30\",\"ACME3_DC_Application41\"", "\"ACME3_DC_Application11\",\"ACME3_DC_Application30\",\"ACME3_DC_Application42\"", "\"ACME3_DC_Application14\",\"ACME3_DC_Application30\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application15\",\"ACME3_DC_Application30\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application16\",\"ACME3_DC_Application31\",\"ACME3_DC_Application43\"", "\"ACME3_DC_Application17\",\"ACME3_DC_Application31\",\"ACME3_DC_Application44\"", "\"ACME3_DC_Application18\",\"ACME3_DC_Application31\",\"ACME3_DC_Application44\"", "\"ACME3_DC_Application19\",\"ACME3_DC_Application31\",\"ACME3_DC_Application45\"", "\"ACME3_DC_Application2\",\"ACME3_DC_Application22\",\"ACME3_DC_Application27\",\"ACME3_DC_Application39\"", "\"ACME3_DC_Application2\",\"ACME3_DC_Application27\",\"ACME3_DC_Application39\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application31\",\"ACME3_DC_Application46\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application32\",\"ACME3_DC_Application34\",\"ACME3_DC_Application47\"", "\"ACME3_DC_Application20\",\"ACME3_DC_Application32\",\"ACME3_DC_Application46\"", "\"ACME3_DC_Application21\",\"ACME3_DC_Application34\",\"ACME3_DC_Application48\"", "\"ACME3_DC_Application21\"", "\"ACME3_DC_Application22\"", "\"ACME3_DC_Application23\"", "\"ACME3_DC_Application24\"", "\"ACME3_DC_Application25\"", "\"ACME3_DC_Application26\",\"ACME3_DC_Application28\",\"ACME3_DC_Application40\",\"ACME3_DC_Application5\"", "\"ACME3_DC_Application26\"", "\"ACME3_DC_Application27\",\"ACME3_DC_Application3\",\"ACME3_DC_Application40\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application4\",\"ACME3_DC_Application40\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application40\",\"ACME3_DC_Application6\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application41\",\"ACME3_DC_Application7\"", "\"ACME3_DC_Application28\",\"ACME3_DC_Application41\",\"ACME3_DC_Application8\"", "\"ACME3_DC_Application29\",\"ACME3_DC_Application41\",\"ACME3_DC_Application9\"", "\"ACME3_DC_Application34\"", "\"ACME3_DC_Application35\"", "\"ACME3_DC_Application36\"", "\"ACME3_DC_Application37\"", "\"ACME3_DC_Application38\"", ""}
	var resIncContextEnv = []string{"UNKNOWN", "DEVELOPMENT", "PRE_PRODUCTION", "DEVELOPMENT", "PRODUCTION", "DEVELOPMENT", "STAGING", "UNKNOWN", "TEST"}
	var resIncTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var resIncHostName = []string{"afkznantre", "awccrrfk56", "awccrrfk57", "awccrrfk58", "awccrrfk59", "awcfvasoxcgfz56", "awicegvopyk56", "awicjpherncc56", "cc-ermcnex", "g6cfprapoffdy56", "inosf58", "ipf7", "ipragre-jqp59", "ipragre-qny65", "j59icefueegz56", "j59icefueegz57", "j59icefueica56", "j59icefuencz56", "j59icefueose56", "j59icefueose57", "j59icefueqaf56", "j59icefueqaf57", "j59icefueqbc56", "j59icefueqbc57", "j59icefueryx51", "j59icefueryx56", "j59icefueryxzba56", "j59icefuexo56", "j59icefuexo57", "j59icefueybt56", "j59icefueybt57", "j59icefuezckl56", "j59icefuezy56", "j59iceoxcgfz56", "j59icerpzjro56", "j59icerpzjro57", "j59icerpzncc50", "j59icerpzncc56", "j59icerpzncc57", "j59icerpzncc58", "j59icerpzncc59", "j59icerpzncc65", "j59icerpznrzc50", "j59icerpznrzc51", "j59icerpznrzc56", "j59icerpznrzc57", "j59icerpznrzc58", "j59icerpznrzc59", "j59icerpznrzn56", "j59icerpznrzn57", "j59icerpzpnu56", "j59icerpzpnu57", "j59icerpzqvep56", "j59icerpzqvep57", "j59icerpzrqo56", "j59icerpzrqo57", "j59icerpzrqo58", "j59icerpzrqo59", "j59icjfrpwzc56", "j59icjfuefcir56", "j59icjfuenq56", "j59icjfuenq57", "j59icjfuewzc56", "j59icjrpzngg56", "j59icjrpzwzc56", "j59igefuerqo56", "j59igerpzjro56", "j59igerpzjro57", "j59igerpzncc52", "j59igerpzncc56", "j59igerpzncc57", "j59igerpzncc58", "j59igerpzncc59", "j59igerpznrz50", "j59igerpznrz51", "j59igerpznrz52", "j59igerpznrz53", "j59igerpznrz54", "j59igerpznrz56", "j59igerpznrz57", "j59igerpznrz58", "j59igerpznrz68", "j59igerpzrqo56", "j59igerpzrqo58", "j59igerpzrqo59", "j59igjrpzwzc56", "jud-jvmpbz-ceq7", "jud-jvmpbz-oqp6", "nmcfnipbecqp56", "nmcfnipbecqp57", "o6cfnicebkgnq56", "o6cfnifrcqz56", "o6cfnifrcva56", "o6cfniibygfdy56", "o6cfniieffdy57", "o6cfninccpbe56", "o6cfnisfpncc56", "o6cfnopragjro56", "o6cnierpbancc56", "o6icynopucncc56", "o6ihynopucncc56", "o78mcejvmncc56", "o78mcejvmncc57", "o78mcejvmzba56", "o78mcejvmzba57", "o78mhejvmncc56", "o7cffnozyttgj57", "o7cfnibz0fbn56", "o7cfnibz0fbn57", "o7cfniepapgk57", "o7cfnipfnszna56", "o7cfnipfuejgf56", "o7cfnippngncc56", "o7cfnipzyttgj53", "o7cfnipzyttgj57", "o7cfnohngfdy56", "o7cfnohngjro56", "o7cfnohngncc56", "o7cfnotbz0qo56", "o7cfnotbz0qo57", "o7cfnotbz0qo58", "o7cfnotgjvncc56", "o7cfnotouq56", "o7cfnotzaqncc56", "o7cfnzqzyttgj57", "o7cfongtqftgj68", "o7cfpbeohfncc58", "o7cfpbeuhzfdy56", "o7cfpbeuhzvak56", "o7qfni9fbntom", "o7qfnibz9ecg56", "o7qfnibz9fbn56", "o7qfnibz9fbnm6", "o7qfnibz9fdy56", "o7qfnohngztg56", "o7qfnotbz8qo56", "o7qfnotbz9qo56", "o7qfnotbz9qo57", "o8ccecebben66", "o8ccecebben67", "o8ccecebben68", "o8ccecebncc68", "o8ccecebncc69", "o8ccecnhben56", "o8ccecrzben56", "o8ccegjben56", "o8ccegjben57", "o8ccegjvncc56", "o8ccegjvncc57", "o8ccenqoben56", "o8ccenqoben57", "o8ccepzfben56", "o8ccepzfncc56", "o8cceqclben56", "o8cceqclncc57", "o8cceqclpqz57", "o8cceqfgben56", "o8cceqfgncc56", "o8cceqfoben56", "o8cceqfoncc56", "o8cceqonben56", "o8cceqonben57", "o8cceqonben58", "o8cceqscben56", "o8cceqyaben56", "o8cceqyancc57", "o8cceqyjben56", "o8cceqypben56", "o8ccerzrncc56", "o8ccerzrncc57", "o8ccexnapfr50", "o8ccjcnhfdy56", "o8ccjfnifsopdq6", "o8ccjfnifsoqve7", "o8ccjfnifsosr56", "o8ccjgcpfpp56", "o8ccjjeyfdy50", "o8ccjjeyfdy57", "o8ccjppnencc56", "o8ccjppnencc57", "o8ccjvasrkz50", "o8ccjvasrkz51", "o8ccjvasrkz56", "o8ccjvasrkz57", "o8ccjvasrkz58", "o8ccrbztfs56", "o8ccrbztfs57", "o8ccrbztfu56", "o8ccrbztfu57", "o8ccrbztfu58", "o8ccrecgrfk56", "o8ccrecgrfk57", "o8ccrfsrfk50", "o8ccrfsrfk51", "o8ccrfsrfk52", "o8ccrfsrfk53", "o8ccrfsrfk56", "o8ccrfsrfk57", "o8ccrfsrfk58", "o8ccrfsrfk59", "o8ccrwhzcfu57", "o8ccrzqzrfk56", "o8ccrzqzrfk57", "o8cfjjeyfdy56", "o8cfjjeyfdy57", "o8cfjjeyfdy59", "o8cfnibz0qgpby6", "o8cfnibz0zd57", "o8cfnibz0zqz57", "o8cfnic3nr56", "o8cfnic3nr57", "o8cfnic3pc56", "o8cfnifzfben56", "o8cfnifzfncc56", "o8cfnigecben56", "o8cfniic3jro56", "o8cfniic3jro57", "o8cfnijngjf57", "o8cfnipxnapfr56", "o8cfnipynben56", "o8cfnjmyqncc57", "o8cfnjmyqncc58", "o8cfnjvmjf57", "o8cfnjvmjf58", "o8cfnjvmjf59", "o8cfnotfgtncc56", "o8cfnotfgtqo56", "o8cfnotvas56", "o8checebben66", "o8checebben67", "o8checebben68", "o8checebncc60", "o8checebncc66", "o8checebncc67", "o8checebncc68", "o8checebncc69", "o8checnhben56", "o8checrzben56", "o8cheenhncc56", "o8chenebflo68", "o8chenebncc66", "o8chenebncc68", "o8chenebncc69", "o8chenqoben56", "o8chenqoben57", "o8chepzfben56", "o8chepzfncc56", "o8cheqanpqz51", "o8cheqanpqz53", "o8cheqanpqz58", "o8cheqclben56", "o8cheqclncc57", "o8cheqclpqz56", "o8cheqfgben56", "o8cheqfoben56", "o8cheqfoncc57", "o8cheqonben56", "o8cheqonben57", "o8cheqscazd56", "o8cheqscben56", "o8cheqyaben56", "o8cheqyjben56", "o8cheqyjncc56", "o8cheqyjncc57", "o8cheqypben56", "o8cheqypncc56", "o8cheqypncc57", "o8cherzrncc56", "o8chjgnfncc56", "o8chjjeyfdy56", "o8chjjeyfdy57", "o8chjqclfdy56", "o8chrbztfs56", "o8chrbztfs57", "o8chrzqzrfk56", "o8cjvasfxyz58", "o8cjvasfxyz59", "o8cjylpcby50", "o8cjylpcby59", "o8cjylppzf50", "o8cjylppzf59", "o8cqenebflo66", "o8cqepynben56", "o8cqeqscazd57", "o8ffprqjc-qo6", "o8ffprqjc-qo7", "o8hfnic3nr56", "o8hfnic3nr57", "o8hfnic3pc56", "o8hfnic3pc57", "o8hfnic3pyqz56", "o8hfnic3pyqz57", "o8hfnifzfben56", "o8hfnifzfncc56", "o8hfniic3fdy56", "o8hfniic3jro56", "o8hfnipzfben56", "o8hfnipzffdy56", "o8icebztfvzncc56", "o8icebztncc50", "o8icebztncc51", "o8icebztncc56", "o8icebztncc57", "o8icebztncc58", "o8icebztncc59", "o8icepngqon56", "o8icepngqon57", "o8icjavegnc56", "o8icjavprnc56", "o8icjavprqo56", "o8icjavprratr56", "o8icjbzincc58", "o8icjcnlyffdy56", "o8icjecgjfncc56", "o8icjecgjfncc57", "o8icjecgjfncc58", "o8icjecgjfncc59", "o8icjecgpzncc56", "o8icjecgpzncc57", "o8icjgrz56", "o8icjjeyjro54", "o8icjnepuefei56", "o8icjnepuejro56", "o8icjnoorffdy57", "o8icjnotbbqtc53", "o8icjnotbbqtp52", "o8icjnotcegt56", "o8icjnotpgkkn50", "o8icjnotpgkkn51", "o8icjpgkkra50", "o8icjpgkkra51", "o8icjpgkkra52", "o8icjpgkkra54", "o8icjpgkkra56", "o8icjpgkqo56", "o8icjpnvff56", "o8icjpoffdy56", "o8icjqnvj57", "o8icjsbqscy56", "o8icjvasacf56", "o8icjvasirnz56", "o8icjvasjro56a", "o8icjvasjro57", "o8icjvasrkqnt56", "o8icjvasrkrqt59", "o8icjwhzcubfg56", "o8icjwhzcubfg57", "o8icyfrpfpna58", "o8ifjjeyjro56", "o8ifjjeyjro57", "o8ihebztncc56", "o8ihebztncc57", "o8ihebztyqnc56", "o8ihezqztoy56", "o8ihjhvcfdy56", "o8ihjhvcjro56", "o8ihjhvcjro57", "o8ihjjeyncc56", "o8ihjnotbbqof50", "o8ihjnotbbqtc57", "o8ihjnotbbqtp56", "o8ihjnotbbqtz58", "o8ihjpbecqp56", "o8ihjylpncc56", "o8ihjylpncc57", "o8iqebztncc56", "o8iqeqonben56", "o8iqjdnncc56", "o8iqjdnqo56", "o8iqjohffdy-g", "o8iqjohffdy56", "o8iqjzyttgj56", "o8qfnibz9hvn56", "o8qfnibz9zqz56", "o8qfniegpben56", "o8qfnigvzncc56", "o8qfnipynben56", "o8qfnipzoncc56", "o8qfnotqrincc56", "o8qfnotvas56", "o8qfnotvas57", "o8qfnotvasyqz56", "o8qjfcecg56", "o8tmcfniperi56", "o8tmcfniperi57", "oc6cfvasrfk56", "oc6cfvasrfk57", "ocvcfvasfwf56", "oemcfninebfym56", "oemcfnipfueqon56", "oemcfnipfuezfp56", "oemcfniplyqqo56", "ofmdfnipfuezfp56", "ofmdfnipfuezfp57", "ofmjepqri57", "ofmqfnipfueqri56", "ofmqfnipfueqri59", "ofmqfniprqjqri56", "ojud-ebire-jjxf", "onopejf050", "onopenc053", "onopenc065", "onopenc066", "onopenc553", "onopenc554", "onopenc561", "onopenc562", "onopenc564", "onopenc565", "onopenc567", "onopenc568", "onopenc569", "onopenc575", "onopenc576", "oqcvasgqztgjl56", "or7qfnisri8frp58", "otmdfnipfuefei56", "otmqfnipfuefei56", "ov6cfvasip56", "ov6cfvasipqo56", "ovcftevaspgk56", "ovcftevaspgk57", "ovcftevaspgk58", "ovcfvasavz96", "ovcfvasegqp57", "ovcfvasfucjro56", "ovcfvasfucncc56", "ovcfvasgfz-yna-serr", "ovcfvasgqz56", "ovcfvasjvaf56", "ovcfvasqqp57", "ovcvasgqztgjl56", "ovqppraqnagqzm7", "ovqppraqnagqzm8", "ovqzmvfn57", "ovvcfqupc59", "pnopenc053", "pnopenc054", "pnopenc060", "pnopenc066", "pnopenc068", "pnopenc069", "pnopenc553", "pnopenc554", "pnopenc561", "pnopenc564", "pnopenc575", "pnopeqo050", "pnopeqo059", "pnopeqo550", "pnopeqo551", "pnopeqo559", "q65icefueegz56", "q65icefueegz57", "q65icefueica56", "q65icefuencz56", "q65icefuencz57", "q65icefueose56", "q65icefueose57", "q65icefueqaf56", "q65icefueqaf57", "q65icefueqbc56", "q65icefueybt56", "q65icefueybt57", "q65icefuezckl56", "q65iceoxcgfz56", "q65icerpzjns57", "q65icerpzjro56", "q65icerpzjro57", "q65icerpzncc50", "q65icerpzncc51", "q65icerpzncc52", "q65icerpzncc53", "q65icerpzncc54", "q65icerpzncc56", "q65icerpzncc57", "q65icerpzncc58", "q65icerpzncc59", "q65icerpzncc65", "q65icerpznczp56", "q65icerpznrzc50", "q65icerpznrzc51", "q65icerpznrzc56", "q65icerpznrzc57", "q65icerpznrzc58", "q65icerpznrzc59", "q65icerpzpnu56", "q65icerpzpnu57", "q65icerpzqvep56", "q65icerpzqvep57", "q65icerpzrqo50", "q65icerpzrqo56", "q65icerpzrqo57", "q65icerpzrqo58", "q65icerpzrqo59", "q65icjfrpwzc56", "q65icjfuefcir56", "q65icjfuenq56", "q65icjfuenq57", "q65icjfueqnvc56", "q65icjfueqnvc57", "q65icjfuewzc56", "q65icjrpzwzc56", "q65igerpzjns56", "q65igerpzjro56", "q65igerpzjro57", "q65igerpzncc56", "q65igerpzncc57", "q65igerpzncc58", "q65igerpzncc59", "q65igerpznczp56", "q65igerpznrz50", "q65igerpznrz51", "q65igerpznrz52", "q65igerpznrz53", "q65igerpznrz54", "q65igerpznrz56", "q65igerpznrz57", "q65igerpznrz58", "q65igerpznrz59", "q65igerpznrz65", "q65igerpznrz66", "q65igerpznrz67", "q65igerpzrqo50", "q65igerpzrqo56", "q65igerpzrqo57", "q65igerpzrqo58", "q65igerpzrqo59", "q65igjrpzwzc56", "qt9bfc7e6iw", "r7cfnipfvtpnc56", "r7cfnipfvtpnc57", "r7cfnipoymqri59", "r7cfnipsygjgf57", "r7cfnipvasuvf54", "r7cfnipvasuvf66", "r7cfnotbz0ncc56", "r7cfnotbz0ncc57", "r7cfnotbz0vag56", "r7cfnotbz0vag57", "r7cfpbecqzmqp57", "r7hfnipfvtpnc56", "r7qfnipfvtpnc56", "r7qfnisri8jf56", "r7qfnotbz8ncc56", "r7qfnotbz9ncc56", "r7qfnotbz9ncc57", "r7qfnotbz9vag56", "r7qfnotbz9vag57", "r7qfnotyzptgj56", "rfkv0", "rfkv1", "rfkv2", "rfkv5", "rfkv6", "rfkv7", "rfkv8", "rfkv9", "ugqp56-221qp154", "ugqp57-221qp154", "zif6", "zif7"}
	var resIncSituation = []string{"apbcvatcebor", "apbcvatcebor:PYBHQ", "cvat", "FE_FAZC BGZN sybbq vf qrgrpgrq:Fcyhax Ebire Nyreg - BGZN sybbq:NOPE-V-PFBCFVAQ", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: gfg-cevprznantre.noenp.arg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: jvmneqthv-hng7.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: NOPEVagreanyVffhvatPN  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5591J Pregvsvpngr rkcverq: NOTFFY  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: CEBF-CEVPVAT  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: EBIRE-ARKGTRA  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: gfg-cevprznantre-jrfgrea  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-anz-qsc-abegurnfg.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-anz-qsc-jrfg.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: hng-zbovyrnccf.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: jvmneqthv-cnlyrffpne  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: jvmneqthvohqtrg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: OCZHNG.NIVFOHQTRG.PBZ  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: qrsnhyg  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC JTNJN5592J Erirefr Cebkl vf abg ehaavat: vevf-hng.nivfohqtrg.pbz  anzr=flfgrz,cevbevgl=zrqvhz", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=1:freireAnzr=JQP-DN-644-NOTPBZD (65.679.1.644: 64266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=4:freireAnzr=JQP-DN-644-Ohqtrg (65.679.1.644: 63266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=6:freireAnzr=Cbfgterf Ragrecevfr Znantre Freire (672.5.5.6: 62286):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=65:freireAnzr=NOTPBZ-DN-Qnyynf-643 (65.679.25.643: 64266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=66:freireAnzr=NIVF-DN-Qnyynf-644 (65.679.25.644: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Freire Qbja:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={5.6,5.7,5.8}:cerivbhfInyhr=5:inyhr=6:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=11.3:inyhr=30.3:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Znk pbaarpg", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=4:freireAnzr=Ohqtrg-DN-Qnyynf-643 (65.679.25.643: 63266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=18.8888888888888:inyhr=32.0:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngv", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=4:freireAnzr=Ohqtrg-DN-Qnyynf-643 (65.679.25.643: 63266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=29.0:inyhr=35.0:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Znk pbaarpgv", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=18.3888888888888:inyhr=34:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Znk p", "FE_FAZC nyregAnzr=Gbgny pbaarpgvbaf nf crepragntr bs znk_pbaarpgvbaf:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={35,30,45}:cerivbhfInyhr=25.0:inyhr=37:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Znk pbaarpgvba = 155", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=1:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = ragrecevf\"", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=3:freireAnzr=Qnyynf-Cebq-643-NOTPBZ (65.679.12.643: 4266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRG nccyvpngvba_anzr = 'RagrecevfrQ", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=4:freireAnzr=Qnyynf-Cebq-643-Ohqtrg (65.679.12.643: 3266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=1:inyhr=6:cerivbhfFgnghf=YBJ:fgnghf=PYRNE:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = ra\"", "\"FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=6:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG 6", "Vf Vqyr? = g", "Hfreanzr = rag\"", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=83:inyhr=5:cerivbhfFgnghf=UVTU:fgnghf=PYRNE:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr fgngr:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=0:inyhr=1:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRYRPG bvq, sbezng_glcr(bvq, AHYY)", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr-va-genafnpgvba fgngr:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=5:inyhr=2:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=Dhrel = FRG FRFFVBA frnepu_cngu", "FE_FAZC nyregAnzr=Pbaarpgvbaf va vqyr-va-genafnpgvba fgngr:freireVQ=65:freireAnzr=NOTPBZ-DN-Qnyynf-643 (65.679.25.643: 64266):guerfubyqInyhr={0,65,60}:cerivbhfInyhr=1:inyhr=5:cerivbhfFgnghf=YBJ:fgnghf=PYRNE:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=JQP-DN-643-Ntrag:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=1.9561405555555555:inyhr=45.6020617555555555:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=JQP-DN-643-Ntrag:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=69.1812767555555555:inyhr=25.3029871555555555:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=", "FE_FAZC nyregAnzr=PCH hgvyvmngvba:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={25,35,30}:cerivbhfInyhr=73.9669835555555555:inyhr=34.7356892555555555:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=4896:inyhr=2447:cerivbhfFgnghf=UVTU:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=2:freireAnzr=Nivf-DN-Qnyynf-643 (65.679.25.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=9870:inyhr=4611:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr \"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=3:freireAnzr=JQP-DN-644-Nivf (65.679.1.644: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=8753:inyhr=1986:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr fv\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=4:freireAnzr=Qnyynf-Cebq-643-Ohqtrg (65.679.12.643: 3266):guerfubyqInyhr={261,364,476}:cerivbhfInyhr=6876:inyhr=6876:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = ohqrpbzc", "Qngnonfr \"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=65:freireAnzr=Qnyynf-Cebq-643-Nivf (65.679.12.643: 2266):guerfubyqInyhr={6579,6773,6081}:cerivbhfInyhr=6437:inyhr=6437:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzc", "Qngnonf\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=66:freireAnzr=NIVF-DN-Qnyynf-644 (65.679.25.644: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=8565:inyhr=1355:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnon\"", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=66:freireAnzr=Qnyynf-Cebq-644-Nivf (65.679.12.644: 2266):guerfubyqInyhr={6579,6773,6081}:cerivbhfInyhr=8717:inyhr=8717:cerivbhfFgnghf=PYRNE:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzc", "Qngnonf\"", "FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=68:freireAnzr=Qnyynf-Cebq-644-NOTPBZ (65.679.12.644: 4266):guerfubyqInyhr={954,067,169}:cerivbhfInyhr=993:inyhr=993:cerivbhfFgnghf=PYRNE:fgnghf=YBJ:qrgnvyrqVasbezngvba=", "\"FE_FAZC nyregAnzr=Qngnonfr fvmr va freire:freireVQ=7:freireAnzr=JQP-DN-643-Nivf (65.679.1.643: 62266):guerfubyqInyhr={0155,1955,3647}:cerivbhfInyhr=2087:inyhr=4299:cerivbhfFgnghf=ZRQVHZ:fgnghf=UVTU:qrgnvyrqVasbezngvba=Qngnonfr anzr = nifrpbzd", "Qngnonfr fvm\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=04.2193040134547:inyhr=47.9330140178220:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-643", "Svyr flfgrz = /\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=7:ntragAnzr=rcnf66-dn-643:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=39.3635304820:inyhr=46.561873670:cerivbhfFgnghf=YBJ:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-643", "Svyr flfgrz = 616.71.43.\"", "\"FE_FAZC nyregAnzr=Zbfg hfrq qvfx crepragntr:ntragVQ=8:ntragAnzr=rcnf66-dn-644:guerfubyqInyhr={35,30,40}:cerivbhfInyhr=23.4100304820:inyhr=46.561873670:cerivbhfFgnghf=PYRNE:fgnghf=ZRQVHZ:qrgnvyrqVasbezngvba=Ntrag anzr = rcnf66-dn-644", "Svyr flfgrz = 616.71.4\"", "FIP:EGRZFgnghf:gvizba:VGZ:7", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):56:54:52.712391-75755651/PN70F", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):66:72:74.739910-75646785/PN70F", "NGZ_:WboAnzr() Ahzore(        ) ZftVq(VRN935R):76:51:64.157393-75646786/PN70F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(275):75755676-54:72:85.450360/PN50F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(VRN935R):75755671-58:55:73.856007/PN50F", "NGZ_:WboAnzr() Ahzore( ) ZftVq(VRN935R):75755671-58:59:51.436116/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):55:51:90.539761-75755671/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):62:53:59.938049-75755668/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):63:67:63.398179-75755669/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):65:75:60.979634-75755668/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):69:73:68.992363-75755667/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):75:78:54.628604-75755669/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):76:86:09.668552-75755658/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP59276) ZftVq(FQJZYT56):76:95:95.840055-75755658/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755686-55:82:96.515217/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755686-55:83:53.319918/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755756-78:84:99.304337/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755757-55:56:02.505872/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755758-66:02:92.466402/PN70F", "NGZ_:WboAnzr(ARGIVRJ) Ahzore(FGP61885) ZftVq(FQJZYT56):75755758-67:58:80.154069/PN70F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(008):75755660-51:70:95.139031/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(018):75755660-51:89:71.183703/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(019):75755660-51:81:85.071212/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(059):75755660-51:55:85.394300/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(061):75755660-51:66:00.917919/PN50F", "NGZ_:WboAnzr(JRFYYRLT) Ahzore(FGP01638) ZftVq(092):75755660-51:77:53.888751/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP01722) ZftVq(QFAW666R):75755652-52:80:60.092573/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP01722) ZftVq(QFAW666R):75755664-51:57:83.593827/PN50F", "NGZ_:WboAnzr(NIVNZFGE) Ahzore(FGP50606) ZftVq(QFAW666R):75646761-51:57:09.440147/PN50F", "NGZ_:WboAnzr(OHQQYVG) Ahzore(FGP63730) ZftVq(QSF5399V):75755679-50:90:55.328696/PN50F", "NGZ_:WboAnzr(OHQQYVG) Ahzore(FGP68216) ZftVq(QSF5399V):75755673-54:97:81.695926/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAG055V):75755676-59:62:88.303018/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAW666R):75755661-60:95:97.533087/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP01862) ZftVq(QFAW666R):75755675-57:02:92.126541/PN50F", "NGZ_:WboAnzr(QFAGZFGE) Ahzore(FGP67686) ZftVq(QFAG055V):75755759-65:91:06.262786/PN50F", "NGZ_:WboAnzr(QFOEZFGE) Ahzore(FGP67673) ZftVq(QFAW666R):75755685-61:59:77.634516/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAG055V):75755676-58:79:50.096956/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAW666R):75755669-66:77:82.870982/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP01724) ZftVq(QFAW666R):75755677-78:95:05.215797/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP50608) ZftVq(QFAW666R):75646760-62:54:63.628294/PN50F", "NGZ_:WboAnzr(QFOGZFGE) Ahzore(FGP67685) ZftVq(QFAW666R):75755674-52:67:88.671237/PN50F", "NGZ_:WboAnzr(QSUFZNI7) Ahzore(FGP01884) ZftVq(NEP5558V):75646785-66:73:69.855100/PN50F", "NGZ_:WboAnzr(UJFZREO7) Ahzore() ZftVq(QSF6434R):75755672-66:59:74.039318/PN50F", "NGZ_:WboAnzr(VZF) Ahzore(FGP51864) ZftVq(QSF7582):55:66:68.892755-75755671/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755672-54:60:58.527614/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755685-69:55:59.537925/PN70F", "NGZ_:WboAnzr(VZFNCCY1) Ahzore() ZftVq(6):75755758-54:60:58.529562/PN70F", "NGZ_:WboAnzr(VZFNCCY4) Ahzore() ZftVq(6):75755759-54:85:59.659029/PN70F", "NGZ_:WboAnzr(VZFNPCO8) Ahzore() ZftVq(QSF6434R):69:57:83.190071-75755658/PN70F", "npe_abdzte_tzdp_fgq", "npe_ajcbeg_j51s_jva", "npe_bssyvar_ezdp_gvi", "npe_bssyvar_tzfj_fgq", "npe_bssyvar_tzfj_pybhq", "npe_cepoyx_tblj_fgq", "npe_ceppch_8agj_fgq", "npe_ceppch_8agj_fgq_Pyq", "npe_ceppchuv_eymp_gvi", "npe_cepzvf_eymj_gvi_wnmmfz", "npe_cepzvf_eyms_gvi_xghpzn", "npe_cepzvf_khkj_fhabar", "npe_cepzvf_khkj_nrz", "npe_cntsvyr_8agj_fgq", "npe_dqrcgu_dzdp_r7cfnotbz0vag6", "npe_dqrcgu_dzds_o7cfnibz0fbn", "npe_dqrcgu_dzds_r7cfnotbz0ncc", "npe_dqrcgu_dzds_r7cfnotbz0vag", "npe_dqrcgu_dzdz_o7cfnibz0fbn56", "npe_dqrcgu_dzdz_r7cfnotbz0vag57", "npe_dqrcgu_dzdz_r7cfnotbz0vag6", "npe_dquv_ezdj_tqfg_erf56ontohqe", "npe_dquv_mzdp_cebq_flfgrz", "npe_dshyy_dzdj_o8icebztncc", "npe_dshyy_dzdp_o8icebztncc", "npe_dzteceo_mzds_cebq", "npe_efgvzybj_egbp_rpc_erf56c05", "npe_efgvzybj_egbp_rpc_erf56cv7", "npe_efgvzybj_egbp_rpc_erf56cv8", "npe_efgvzybj_egbp_rpc_erf56cv9", "npe_efgvzybj_egbp_rpc_erf56l05", "npe_efgvzybj_egbp_rpc_erf56lv8", "npe_efgvzybj_egbp_rpc_erf56o05", "npe_efgvzybj_egbp_rpc_erf56ov7", "npe_efgvzybj_egbp_rpc_erf56ov8", "npe_efgvzybj_egbp_rpc_erf56ov9", "npe_efgvzybj_egbp_tqfc_erf56cn6", "npe_efgvzybj_egbp_tqfc_erf56cn8", "npe_efgvzybj_egbp_tqfc_erf56cnr", "npe_efgvzybj_egbp_tqfc_erf56cnt", "npe_efgvzybj_egbp_tqfc_erf56lnr", "npe_efgvzybj_egbp_tqfc_erf56lnt", "npe_efgvzybj_egbp_tqfc_erf56on8", "npe_efgvzybj_egbp_tqfc_erf56onr", "npe_efgvzybj_egbp_tqfc_erf56ont", "npe_erderfc_tlaj_jf", "npe_feifgng_tlaj_jf", "npe_feifgng_tlap_jf", "npe_ffyntr_t9fp_ffypreg", "npe_fieepie_tizj_rfk", "npe_fiegfze_tizj_rfk", "npe_fiepch_tizj_rfk", "npe_fieqfse_tizj_rfk", "npe_fiezrz_tizj_rfk", "npe_fiezrz_tizp_rfk", "npe_fip_8agp_fgq", "npe_fip_8agp_fgq_Pyq", "npe_fip_8agp_fgvfip", "npe_fjnccep_fhkj_fgq", "npe_fjnccep_khkj_fgq", "npe_fvgrqja_tugj_uggc", "npe_fvgrsy_tugj_uggc", "npe_fvgrsy_tugp_uggc", "npe_gen_eymj_fgq", "npe_gen_eymj_gvi", "npe_gen_fhkj_fgq", "npe_gen_khkj_fgq", "npe_gffgng_temj_ben", "npe_gfhody_8rkj_fgq", "npe_jpegrkc_8d2j_fgq_i7", "npe_mbz_eymp_fgq_Pyq", "npe_mbz_fhkp_fgq", "npe_mbz_khkp_fgq", "npe_pch_eymj_fgq_pyq", "npe_pch_eymp_cebq", "npe_pch_eymp_fgq", "npe_pch_fhkp_fgq", "npe_pch_khkp_fgq", "npe_pheqcg_tzdp_fgq", "npe_pheqcg_tzds_fgq", "npe_pufgng_ezdp_grfg_jvmzbqgfg", "npe_qfc_8agj_fgq7i7", "npe_qfc_8agp_fgq", "npe_qfc_8agp_fgq_Pyq", "npe_qfc_8ags_fgq", "npe_qfc_8ags_fgq_Pyq", "npe_qobssya_8bdp_fgq", "npe_qofgng_tblp_fgq", "npe_qosvyr_temj_ben", "npe_qovanpg_8bdp_fgq", "npe_qovanpg_temj_fgqdn", "npe_qovanpg_temp_fgq", "npe_qovanpg_temp_fgqdn_rouf", "npe_qovapfg_8bdp_fgq", "npe_qoyef_8rkj_fgq", "npe_qsfgng_temp_fgq", "npe_qvfxfc_temj_ben", "npe_qvfxfc_temp_ben", "npe_qvfxfc_temz_bendn", "npe_reebe_euqp_gvi", "npe_reeybt_fhyj_nonphf", "npe_reeybt_fhyj_nonphf7", "npe_sergofc_temj_ben", "npe_sergofc_temj_ben_rouf", "npe_sergofc_temj_bendn", "npe_sergofc_temj_bendn_rouf", "npe_sergofc_temj_o8ffprqjc", "npe_sergofc_temp_ben", "npe_sergofc_temp_o8ffprqjc", "npe_sergofc_temz_bendn", "npe_sergofc_temz_bendn_rouf", "npe_sff_eymj_fgq", "npe_sff_eymj_fgq_cebq", "npe_sff_eymj_fgq_Pyq", "npe_sff_eymj_gvi_cebq", "npe_sff_eymj_zdz", "npe_sff_eymp_fgq", "npe_sff_eymp_fgq_Pyq", "npe_sff_eymp_gvi_cebq", "npe_sff_eymp_ncc_grfg", "npe_sff_eymp_pyqsvyrcrez", "npe_sff_eyms_fgq", "npe_sff_eyms_fgq_Pyq", "npe_sff_eyms_gvi_cebq", "npe_sff_fhkj_fgq", "npe_sff_fhkj_ROUF", "npe_sff_fhkj_ROUF_gfz", "npe_sff_fhkj_ROUF_grfg", "npe_sff_fhkp_fgq", "npe_sff_fhkp_ROUF", "npe_sff_fhks_fgq", "npe_sff_fhks_ROUF", "npe_sff_khkj_fgq", "npe_sff_khkj_gfzfgt", "npe_sff_khkj_o7cfnotouq56", "npe_sff_khkj_o7qfni9fbntom", "npe_sff_khkj_o7qfnibz9fbn56", "npe_sff_khkj_QN", "npe_sff_khkj_zdz", "npe_sff_khkp_fgq", "npe_sff_khkp_o7qfni9fbntom", "npe_sff_khkp_o7qfnibz9fbn56", "npe_sff_khkp_pfg", "npe_sff_khkp_pfg9", "npe_sff_khkp_zdz", "npe_sff_khks_fgq", "npe_sff_khks_o7qfnibz9fbn56", "npe_sff_khks_zdz", "npe_sofcnpr_temj_bendn", "npe_uggcfg_fugj_uggc", "npe_uggcfg_tugj_uggc", "npe_uogg_eymp_gvi", "npe_uozd_eymp_gvi", "npe_vaf_eymp_fgq", "npe_vaf_fhkp_fgq", "npe_vafgnpg_temj_bendn", "npe_vafgnpg_temp_ben", "npe_vafgnpg_temp_bendn_rouf", "npe_wbosnvy_tbdp_zffdy_cebq", "npe_ybtben_temj_bendn7_rouf", "npe_ybtben_temz_bendn9_rouf", "npe_ybtserr_tblp_fgq", "npe_znkpua_ezdp_cebq_jvmncc56", "npe_znkpua_ezdp_cebq_jvmncc57", "npe_znkpua_ezds_cebq_jvmncc56", "npe_znkpua_ezds_cebq_jvmncc57", "NPF Flagurgvp Gerr Nhgbzngvba Rirag_5", "VGZ:75646786-56"}
	var resIncStatus = []string{"CANCELLED", "CLOSED", "INPROG", "PENDING", "QUEUED", "REJECTED", "RESOLVCONF", "RESOLVED", "SLAHOLD"}
	var resIncTicketClass = []string{"Application Down", "Backup Missed/Failed", "CPU High Issue", "Database Handler", "Disk Storage Issues", "High Memory & Page File Usage", "ITM/Other Agent", "Job Abends", "MQ Handler", "Non-os Windows Disk Full", "Other Automata", "Server Unavailable", "Service in Alert State", "Swap Space Issue", "Table Space Handler", "Unclassified Actionable", "Windows OS Disk Full", "Windows Service Handler", "Zombie Processes"}

	flag := true
	for gbSize < 200 {
		csvData7resinc = nil
		csvData7resinc = append(csvData7resinc, resIncStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand actionable
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncActionable[rand.Intn(len(resIncActionable))])

			//rand assignment_group
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncAssignGroup[rand.Intn(len(resIncAssignGroup))])

			//rand autogenerated
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncAutoGen[rand.Intn(len(resIncAutoGen))])

			//rand BACID
			min := 0200000
			max := 9999999
			randClientID := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand context.application
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "["+resIncContextApp[rand.Intn(len(resIncContextApp))]+"]")

			//rand context.environemnt
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\""+resIncContextEnv[rand.Intn(len(resIncContextEnv))]+"\"]")

			//rand context.manage
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\"IBM\"]")

			//rand context.team
			min = 1
			max = 9
			randTeam := rand.Intn(max-min+1) + min
			//fmt.Println(randTeam)
			tmpSlice = append(tmpSlice, "[\"acme sales demo"+strconv.Itoa(randTeam)+"\"]")

			//rand TENANT_ID part1
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = resIncTenantID[rand.Intn(len(resIncTenantID))]
			}

			//rand correlation_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 0000000
			max = 9999999
			seqno := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "IBM-"+string(TenantIDRune)+"-BAM00ACME3-"+strconv.Itoa(seqno))

			//rand created
			ranCreated, ranTimeStampCreated, _ := randate()
			tmpSlice = append(tmpSlice, ranCreated)

			//rand data_source
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "ACME Corp3")

			//rand datacenter
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "IBM")

			//rand description
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand hostname
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncHostName[rand.Intn(len(resIncHostName))])

			//rand last_refresh
			min = 256
			max = 512
			randRef := rand.Intn(max-min+1) + min
			timeStampRef := ranTimeStampCreated.Add(time.Duration(randRef) * time.Minute)
			timeStampA := timeStampRef.Format("Jan 02, 2006")
			timeStampB := timeStampRef.Format("15:04:05")
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")

			//rand modified_dttm
			min = 24
			max = 128
			randMod := rand.Intn(max-min+1) + min
			timeStampMod := ranTimeStampCreated.Add(time.Duration(randMod) * time.Minute)
			timeStampA = timeStampMod.Format("Jan 02, 2006")
			timeStampB = timeStampMod.Format("15:04:05")
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")

			//rand number
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 000000000
			max = 999999999
			number := rand.Intn(max-min+1) + min
			//fmt.Println(number)
			tmpSlice = append(tmpSlice, "IN"+strconv.Itoa(number))

			//rand priority
			min = 1
			max = 5
			randPriority := rand.Intn(max-min+1) + min
			//fmt.Println(randPriority)
			tmpSlice = append(tmpSlice, strconv.Itoa(randPriority))

			//rand provider_account
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand resolution
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand resolved
			min = 128
			max = 256
			randRes := rand.Intn(max-min+1) + min
			timeStampRes := ranTimeStampCreated.Add(time.Duration(randRes) * time.Minute)
			timeStampA = timeStampRes.Format("Jan 02, 2006")
			timeStampB = timeStampRes.Format("15:04:05")
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")

			//rand situation
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncSituation[rand.Intn(len(resIncSituation))])

			//rand SOURCE_TYPE
			tmpSlice = append(tmpSlice, "DataCenter")

			//rand status
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncStatus[rand.Intn(len(resIncStatus))])

			//rand symptom
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand TENANT_ID part2
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//rand ticket_class
			tmpSlice = append(tmpSlice, "INCIDENT")

			//rand ticket_classification
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, resIncTicketClass[rand.Intn(len(resIncTicketClass))])

			//create the final slice
			csvData7resinc = append(csvData7resinc, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData7resinc)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData7resinc)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/resource_incidents_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData7resinc)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "resource_incidents_temp.csv") + " > " + filepath.Join(currentPath, "resource_incidents_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted asset")
		}

		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData7resinc)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "resource_incidents_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length for resInc", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "resource_incidents", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-resource_incidents.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-resource_incidents.json", head, eUser, ePassword, elasticClusterIP, "resource_incidents")

			//cleanup the file
			cleanup(currentPath, "finalOutput-resource_incidents.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-resource_incidents.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-resource_incidents.json", head, eUser, ePassword, elasticClusterIP, "resource_incidents")

		//cleanup the file
		cleanup(currentPath, "finalOutput-resource_incidents.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/resource_incidents/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			//fmt.Println("Response for increasing limit", response)

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/resource_incidents/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("resource_incidents index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		//time.Sleep(5 * time.Second)

		////
		/////
		//////
	}

	defer wg.Done()
}

func dataSetAsset(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData6asset = [][]string{
		{"SOURCE_TYPE", "UID", "D_CLIENT_ID", "TENANT_ID", "BACID", "SYSTEM_ID", "MANUFACTURER", "MODEL", "MACHINETYPE", "SERIALNUMBER", "SERIAL5", "ASSETOWNER", "ASSETCLS", "ASSETTAG", "ASSETTYPE", "OWNER", "PROCESSOR_TYPE", "CLASSIFICATION", "ASSET_LIFECYCLESTATE", "CATEGORY", "CI_LIFECYCLESTATE", "HW_SEGMENT", "DATASOURCE", "INSTALL_DATE", "LASTMODIFIEDBY", "LASTMODIFIEDTIME", "PLATFORM", "PURPOSE", "ISOCNTRY", "CDIR_ID", "CHIP_ID", "CHIP_CLIENT_UID", "CDIR_CLIENT_UID", "CNDB_ID", "PRIMNAME", "CDIR_CUSTNAME", "CHIP_CUSTNAME", "FQHN", "OSNAME", "OSNAME2", "OSVERSION", "OSRELEASE", "OSSUBVERSION", "DELETED", "VALID_FLAG", "CMDB_FLAG", "HOSTNAME", "ASSET_ID", "ATP_UID", "BRAVO_UID", "SRM_UID", "ECM_ORPHAN_UID", "SYSTEM_COMPLAINCE_SCORE_SYSTEM_ID", "NEVER_MOVED_UID", "EVENT_ID", "UNRECONCILED_CI", "CHIP_CLIENT_ECM_FLAG", "CHIP_CLIENT_DELETED", "CHIP_CLIENT_STATUS", "EVENT_UID", "SEA_UID", "IEM_UID", "IP_ADDRESS", "TFLAG", "OSPROVIDER", "ECM_SECTOR", "HOSTSYSTEM_ID", "HOSTSYSTEM_TYPE", "RECON_UID", "OSMODLEVEL", "OS_PATCH_LEVEL", "SSAE16", "GIAM_UID", "IPSOFT_FLAG", "ACC_INFRA_UID", "ACC_BILL_UID", "FP", "FP_TYPE", "FP_ATP", "FP_BRAVO", "FP_GIAMA", "FP_IEM", "ATP_MATCH", "BRAVO_MATCH", "GIAMA_MATCH", "SEA_MATCH", "NETCOOL_MATCH", "IEM_MATCH", "SRM_MATCH", "FP_SEA", "FP_NETCOOL", "FP_SRM", "USE_ALIAS_HOSTNAME", "ALIAS_HOSTNAME", "ATP_HOSTNAME", "BRAVO_HOSTNAME", "SRM_HOSTNAME", "GIAMA_HOSTNAME", "SEA_HOSTNAME", "NETCOOL_HOSTNAME", "ACC_INFRA_HOSTNAME", "ACC_BILL_HOSTNAME", "IEM_HOSTNAME", "ECM_HOSTNAME", "ECM_ORPHAN_HOSTNAME", "STATE_DEVIATION", "EOL", "CREATE_DATE", "TOTAL_USERS", "USERS_EXTERNAL", "EOL_UID", "SB_DATE", "SB_FLAG", "SB_UID", "VIRTUAL_FLAG", "OSTYPE", "ECM_MATCH", "ECM_ORPHAN_MATCH", "ACCOUNT_INFRA_MATCH", "ACCOUNT_BILLING_MATCH", "FP_ECM_ORPHAN", "FP_ECM", "FP_ACC_INFRA", "FP_ACC_BILL", "PARENT_UID", "LPAR_NAME", "ENVIRONMENT", "SERVER_TYPE", "PRODDATE", "ATP_SUID", "ATP_SMATCH", "ATP_SERIAL", "SEA_SUID", "SEA_SMATCH", "SEA_SERIAL", "SRM_SUID", "SRM_SMATCH", "SRM_SERIAL", "BRAVO_SUID", "BRAVO_SMATCH", "BRAVO_SERIAL", "P_CI_LIFECYCLESTATE", "NETCOOL_LDS_UID", "FP_NETCOOL_LDS", "NETCOOL_LDS_HOSTNAME", "NETCOOL_LDS_MATCH", "MAJOR_BUS_PROC", "IPC_SERVER_ID", "IPC_SERVER_HOSTNAME", "CURRENCY_DESIGNATOR", "DYNAMIC_AUTOMATION_FLAG", "DYNAMIC_AUTOMATION_BETA_FLAG", "MATCH_NETCOOL_LDS_DAFLAG", "MATCH_IPC_SERVER_DAFLAG", "NETCOOL_LDS_ALIAS_UID", "NETCOOL_ALIAS_HOSTNAME", "SECURITY_VIOLATIONS", "EVENTS", "INCIDENTS", "SERVICE_REQUESTS", "PROBLEMS", "CHANGES", "INTERNET_ACC_FLAG", "DR_FLAG", "SECURITY_CLASSIFICATION", "IBM_CBN_INTERVAL", "PRIV_ID_INTERVAL", "HC_AUTO_INTERV", "HCLASS", "APPL_SLA", "SUBSYS_COUNT", "IBM_QEV_INTERVAL", "OS_SEGMENT", "PATCH_SCORE", "HC_SCORE", "COMPLIANCE_RISK_SCORE", "CPU_UTILIZATION", "MEMORY_UTILIZATION", "MEMORY_SIZE", "PURPOSE2", "CDI_CREATE_DTTM", "TOOL_SOURCE", "S_E_FOR_A_T", "CAPDATE", "HWAGE", "PLATFORM_TYPE_ID", "NODE_ID", "LAST_REFRESH_DTTM"},
	}

	//sample strings asset
	var assetStartString = []string{"SOURCE_TYPE", "UID", "D_CLIENT_ID", "TENANT_ID", "BACID", "SYSTEM_ID", "MANUFACTURER", "MODEL", "MACHINETYPE", "SERIALNUMBER", "SERIAL5", "ASSETOWNER", "ASSETCLS", "ASSETTAG", "ASSETTYPE", "OWNER", "PROCESSOR_TYPE", "CLASSIFICATION", "ASSET_LIFECYCLESTATE", "CATEGORY", "CI_LIFECYCLESTATE", "HW_SEGMENT", "DATASOURCE", "INSTALL_DATE", "LASTMODIFIEDBY", "LASTMODIFIEDTIME", "PLATFORM", "PURPOSE", "ISOCNTRY", "CDIR_ID", "CHIP_ID", "CHIP_CLIENT_UID", "CDIR_CLIENT_UID", "CNDB_ID", "PRIMNAME", "CDIR_CUSTNAME", "CHIP_CUSTNAME", "FQHN", "OSNAME", "OSNAME2", "OSVERSION", "OSRELEASE", "OSSUBVERSION", "DELETED", "VALID_FLAG", "CMDB_FLAG", "HOSTNAME", "ASSET_ID", "ATP_UID", "BRAVO_UID", "SRM_UID", "ECM_ORPHAN_UID", "SYSTEM_COMPLAINCE_SCORE_SYSTEM_ID", "NEVER_MOVED_UID", "EVENT_ID", "UNRECONCILED_CI", "CHIP_CLIENT_ECM_FLAG", "CHIP_CLIENT_DELETED", "CHIP_CLIENT_STATUS", "EVENT_UID", "SEA_UID", "IEM_UID", "IP_ADDRESS", "TFLAG", "OSPROVIDER", "ECM_SECTOR", "HOSTSYSTEM_ID", "HOSTSYSTEM_TYPE", "RECON_UID", "OSMODLEVEL", "OS_PATCH_LEVEL", "SSAE16", "GIAM_UID", "IPSOFT_FLAG", "ACC_INFRA_UID", "ACC_BILL_UID", "FP", "FP_TYPE", "FP_ATP", "FP_BRAVO", "FP_GIAMA", "FP_IEM", "ATP_MATCH", "BRAVO_MATCH", "GIAMA_MATCH", "SEA_MATCH", "NETCOOL_MATCH", "IEM_MATCH", "SRM_MATCH", "FP_SEA", "FP_NETCOOL", "FP_SRM", "USE_ALIAS_HOSTNAME", "ALIAS_HOSTNAME", "ATP_HOSTNAME", "BRAVO_HOSTNAME", "SRM_HOSTNAME", "GIAMA_HOSTNAME", "SEA_HOSTNAME", "NETCOOL_HOSTNAME", "ACC_INFRA_HOSTNAME", "ACC_BILL_HOSTNAME", "IEM_HOSTNAME", "ECM_HOSTNAME", "ECM_ORPHAN_HOSTNAME", "STATE_DEVIATION", "EOL", "CREATE_DATE", "TOTAL_USERS", "USERS_EXTERNAL", "EOL_UID", "SB_DATE", "SB_FLAG", "SB_UID", "VIRTUAL_FLAG", "OSTYPE", "ECM_MATCH", "ECM_ORPHAN_MATCH", "ACCOUNT_INFRA_MATCH", "ACCOUNT_BILLING_MATCH", "FP_ECM_ORPHAN", "FP_ECM", "FP_ACC_INFRA", "FP_ACC_BILL", "PARENT_UID", "LPAR_NAME", "ENVIRONMENT", "SERVER_TYPE", "PRODDATE", "ATP_SUID", "ATP_SMATCH", "ATP_SERIAL", "SEA_SUID", "SEA_SMATCH", "SEA_SERIAL", "SRM_SUID", "SRM_SMATCH", "SRM_SERIAL", "BRAVO_SUID", "BRAVO_SMATCH", "BRAVO_SERIAL", "P_CI_LIFECYCLESTATE", "NETCOOL_LDS_UID", "FP_NETCOOL_LDS", "NETCOOL_LDS_HOSTNAME", "NETCOOL_LDS_MATCH", "MAJOR_BUS_PROC", "IPC_SERVER_ID", "IPC_SERVER_HOSTNAME", "CURRENCY_DESIGNATOR", "DYNAMIC_AUTOMATION_FLAG", "DYNAMIC_AUTOMATION_BETA_FLAG", "MATCH_NETCOOL_LDS_DAFLAG", "MATCH_IPC_SERVER_DAFLAG", "NETCOOL_LDS_ALIAS_UID", "NETCOOL_ALIAS_HOSTNAME", "SECURITY_VIOLATIONS", "EVENTS", "INCIDENTS", "SERVICE_REQUESTS", "PROBLEMS", "CHANGES", "INTERNET_ACC_FLAG", "DR_FLAG", "SECURITY_CLASSIFICATION", "IBM_CBN_INTERVAL", "PRIV_ID_INTERVAL", "HC_AUTO_INTERV", "HCLASS", "APPL_SLA", "SUBSYS_COUNT", "IBM_QEV_INTERVAL", "OS_SEGMENT", "PATCH_SCORE", "HC_SCORE", "COMPLIANCE_RISK_SCORE", "CPU_UTILIZATION", "MEMORY_UTILIZATION", "MEMORY_SIZE", "PURPOSE2", "CDI_CREATE_DTTM", "TOOL_SOURCE", "S_E_FOR_A_T", "CAPDATE", "HWAGE", "PLATFORM_TYPE_ID", "NODE_ID", "LAST_REFRESH_DTTM"}
	var assetTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var assetSystemID = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var assetManufacturer = []string{"CISCO", "DELL", "DUMMYGEN", "HPE", "IBM", "SUPERMICRO"}
	var assetModel = []string{"1950", "48PS-S", "AC1", "AG-A-K9", "AG-E-K9", "BL460C G6", "BL460C G7", "BL460C G8", "DL320E G8", "DL380 G5", "DL380P G8", "DUMM", "HSEC/K9", "I-A-K9", "K9", "N-E-K9", "R520", "R810", "S-48LPD-L", "SLBM", "X-48FPS-L"}
	var assetMachineType = []string{"7143", "AIR-AP1252", "AIRCAP2602", "AIRLAP1131", "AIRLAP1141", "AIRLAP1252", "CISCO2811", "CISCO2921", "CISCO3945", "DUMM", "POWEREDGE", "PROLIANT", "SLBM", "WS-C2960", "WS-C3750"}
	var assetSerialNumber = []string{"1G8CXF6", "2WFTFE6", "3WFTFE6", "4JOKIO6", "51QU236", "7HK38056IA", "7Z787959D3", "FTU967R614", "O8WDDI6", "P374HNR52N85109", "P374HNR58N75102", "P374HNR90O56512", "P374HNR90O56619", "P374HNR91O65794", "P374HNR98O55066", "P374HNS65O65584", "P374HNS65O65598", "P374HNS65O65604", "PM860584PL", "PM860584Q0", "PM860584Q3", "PM860584Q7", "PMW47450AA", "PMW57052U0", "PMW57052U9", "PMW57052UP", "PMW57052UU", "PMW6895PG2", "PMW6895PG3", "PMW6895SUG", "PMW69457GW", "PMW69457GX", "PMW7605139", "PMW76062OL", "PMW766645E", "PMW7675Y0A", "PMW7675Y0T", "PMW7675Y0X", "PMW771525X", "PNG6588AWM3", "SGK699146PQ", "SPJ6151Y5AS", "SPJ6386P6N5", "SPJ6392N1ZU", "SPJ6773M597", "SPJ6773M598", "SPJ7559ANE3", "SPM6065D5IQ", "SPM6854250U", "SPM6891D6E1", "SPU6178I1AG", "STY6265F55M", "STY6265F577", "TRAJCC"}
	var assetContantSlice = []string{"", "", "", "", "UPGRADE", "", "", "", "IN_USE", "CPU", "PRODUCTION", "", "ACME", "26/04/2017", "", "00:00.0", "", "", "GB", "CDI-0009999123", "1000002", "9999123", "9999123", "9999123", "Acme", "Acme", "Acme", "ytjq6cwzc5557", "RED HAT ENTERPRISE LINUX", "LINUX", "6", "6", "6", "N", "", "Y", "ytjq6cwzc5557", "", "", "", "", "", "", "", "", "N", "", "", "", "0", "0", "0", "", "", "REDHAT", "", "", "PHYCOMPSYS", "0", "", "", "N", "", "", "", "", "Y", "UNKNOWN", "N", "N", "N", "N", "", "", "", "", "", "", "", "N", "N", "N", "", "", "", "", "", "", "", "", "", "", "", "", "", "N", "N", "26/04/2017", "", "", "", "", "", "", "", "LINUX", "", "", "RULE 1", "RULE 1", "Y", "Y", "Y", "N", "", "", "LINUX", "PRODUCTION", "26/09/2017", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "N", "", "RULE 1", "", "", "", "", "", "", "N", "", "", "", "", "", "", "", "", "", "N", "N", "CUSTOMER CONFIDENTIAL", "0", "0", "0", "ISCD", "", "1", "3", "REDHAT 6", "", "", "", "", "", "", "", "09:01.3", "SESDR", "N", "", "", "1", "", "09:01.3"}

	flag := true
	for gbSize < 200 {
		csvData6asset = nil
		csvData6asset = append(csvData6asset, assetStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand SOURCE_TYPE
			tmpSlice = append(tmpSlice, "DataCenter")

			//rand UID
			rand.Seed(time.Now().UnixNano())
			min := 10000000
			max := 99999999
			randID := rand.Intn(max-min+1) + min
			//fmt.Println(rand.Intn(max-min+1) + min)
			tmpSlice = append(tmpSlice, strconv.Itoa(randID))

			//rand D_CLIENT_ID
			min = 9999123
			max = 9999999
			randClientID := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, strconv.Itoa(randClientID))

			//rand TENANT_ID
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = assetTenantID[rand.Intn(len(assetTenantID))]
			}
			//fmt.Println(TenantIDRune)
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//rand BACID
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand SYSTEM_ID
			SystemIDRune := make([]rune, 1)
			for i := range SystemIDRune {
				SystemIDRune[i] = assetSystemID[rand.Intn(len(assetSystemID))]
			}
			min = 1000000
			max = 9999999
			randNumber := rand.Intn(max-min+1) + min
			tmpSlice = append(tmpSlice, string(SystemIDRune)+strconv.Itoa(randNumber))

			//rand MANUFACTURER
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, assetManufacturer[rand.Intn(len(assetManufacturer))])

			//rand MODEL
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, assetModel[rand.Intn(len(assetModel))])

			//rand MACHINETYPE
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, assetMachineType[rand.Intn(len(assetMachineType))])

			//rand SERIALNUMBER
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, assetSerialNumber[rand.Intn(len(assetSerialNumber))])

			for _, row1 := range assetContantSlice {
				tmpSlice = append(tmpSlice, row1)
			}

			//tmpStruct = append(tmpStruct, tmpSlice)

			//create the final slice
			csvData6asset = append(csvData6asset, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData6asset)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData6asset)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/ACME_asset_data_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData6asset)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "ACME_asset_data_temp.csv") + " > " + filepath.Join(currentPath, "ACME_asset_data_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted asset")
		}

		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData6asset)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "ACME_asset_data_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length for asset", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "asset", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-asset.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-asset.json", head, eUser, ePassword, elasticClusterIP, "asset")

			//cleanup the file
			cleanup(currentPath, "finalOutput-asset.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-asset.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-asset.json", head, eUser, ePassword, elasticClusterIP, "asset")

		//cleanup the file
		cleanup(currentPath, "finalOutput-asset.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/asset/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			//fmt.Println("Response for increasing limit", response)

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/asset/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("Asset index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		time.Sleep(5 * time.Second)

		////
		/////
		//////
	}

	defer wg.Done()
}

func dataSetChange(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData5chg = [][]string{
		{"SOURCE_TYPE", "ID", "D_CLIENT_ID", "TENANT_ID", "BACID", "CUSTOMER_NAME", "CHANGE_ID", "S_DATA_SOURCE", "REQUEST_DTTM", "PLAN_ST_DTTM", "ACT_ST_DTTM", "ACT_FINISH_DTTM", "SCHEDULED_DTTM", "SCHED_FINISH_DTTM", "CLOSED_DTTM", "CATEGORY_CD", "CHANGE_TYPE_CD", "STATUS_CD", "RISK_CODE_CD", "ASSIGNEE_GROUP_CD", "CI_DESC", "CHNG_ABSTRCT_DESC", "CHANGE_DESC", "CMPLTN_CODE_CD", "AGE", "AGE_BUCKET", "SEVERITY", "TICKET_CLASS", "DEFERENCE_WORK_STARTED", "DEFERENCE_WORK_END", "REGION", "COMPANY", "PARENT", "EXCEPTION_JUSTIFICATION", "EXCEPTION_REASON", "CREATED", "CREATED_DAY", "CREATED_TIME", "CREATED_MONTH", "CREATED_WEEK", "CLOSED_DAY", "CLOSED_TIME", "CLOSED_MONTH", "CLOSED_WEEK", "LEAD_TIME", "PERSON_STREET_ADDRESS", "PERSON_ADDRESSLINE2", "PERSON_ADDRESSLINE3", "PERSON_LOCATION", "PERSON_REGION", "PERSON_COUNTRY", "PERSON_SITE_ID", "PERSON_BUILDING_ID", "PERSON_CITY", "PERSON_DEPARTMENT", "LAST_REFRESH_DTTM", "ACTUALENDDATE", "ACTUALSTARTDATE", "ASSIGNEE", "ASSIGNEEGROUP", "CHANGEID", "CLOSURECODE", "COMPLETEDDATE", "ENVIRONMENT", "MODIFIEDDATE", "PRIORITY", "PRODUCTCATEGORIZATIONTIER1", "PRODUCTCATEGORIZATIONTIER2", "PRODUCTCATEGORIZATIONTIER3", "RELATEDCINAME", "RELATEDINCIDENTID", "RFCDATE", "RISKLEVEL", "SCHEDULEDENDDATE", "SCHEDULEDSTARTDATE", "STATUS", "SUBMITDATE", "SUMMARY", "TARGETDATE", "YEAR_MONTH_SCH_DTTM", "HOUR_SCH_DTTM", "DAY_NAME_SCH_DTTM", "YEAR_WEEK_SCH_DTTM", "YEAR_MONTH_ACTST_DTTM", "HOUR_ACTST_DTTM", "DAY_NAME_ACTST_DTTM", "YEAR_MONTH_ACTFT_DTTM", "HOUR_ACTFT_DTTM", "DAY_NAME_ACTFT_DTTM", "YEAR_WEEK_ACTFT_DTTM", "YEAR_MONTH_SCHFT_DTTM", "HOUR_SCHFT_DTTM", "DAY_NAME_SCHFT_DTTM", "COMPARE_START", "COMPARE_FINISH", "AGEBIN", "TTR_DAYS", "TTR_BIN", "FASTPATH_REASON_CD", "ITDCLOSURECODE", "SUCCESSVSFAILED", "ITDEXCEPTION", "PMCHGTYPE", "CANCELLED_COUNT", "CLOSED_COUNT", "SUCCESS_COUNT", "SUCCESSFULPERCENT", "FAIL_COUNT", "FAILEDPERCENT", "EMERGENCY_COUNT", "EMERGENCYPERCENT", "UNAUTH_COUNT", "UNAUTHPERCENT", "APPROVAL_CD", "APPROVED_COUNT", "APPROVEDPERCENT", "REJECTED_COUNT", "REJECTEDPERCENT", "EXCEPTIONPERCENT", "BACKLOG", "COUNT", "YEAR_MONTH_CREATED", "YEAR_MONTH_CLOSED"},
	}

	//sample strings change
	var chgStartString = []string{"SOURCE_TYPE", "ID", "D_CLIENT_ID", "TENANT_ID", "BACID", "CUSTOMER_NAME", "CHANGE_ID", "S_DATA_SOURCE", "REQUEST_DTTM", "PLAN_ST_DTTM", "ACT_ST_DTTM", "ACT_FINISH_DTTM", "SCHEDULED_DTTM", "SCHED_FINISH_DTTM", "CLOSED_DTTM", "CATEGORY_CD", "CHANGE_TYPE_CD", "STATUS_CD", "RISK_CODE_CD", "ASSIGNEE_GROUP_CD", "CI_DESC", "CHNG_ABSTRCT_DESC", "CHANGE_DESC", "CMPLTN_CODE_CD", "AGE", "AGE_BUCKET", "SEVERITY", "TICKET_CLASS", "DEFERENCE_WORK_STARTED", "DEFERENCE_WORK_END", "REGION", "COMPANY", "PARENT", "EXCEPTION_JUSTIFICATION", "EXCEPTION_REASON", "CREATED", "CREATED_DAY", "CREATED_TIME", "CREATED_MONTH", "CREATED_WEEK", "CLOSED_DAY", "CLOSED_TIME", "CLOSED_MONTH", "CLOSED_WEEK", "LEAD_TIME", "PERSON_STREET_ADDRESS", "PERSON_ADDRESSLINE2", "PERSON_ADDRESSLINE3", "PERSON_LOCATION", "PERSON_REGION", "PERSON_COUNTRY", "PERSON_SITE_ID", "PERSON_BUILDING_ID", "PERSON_CITY", "PERSON_DEPARTMENT", "LAST_REFRESH_DTTM", "ACTUALENDDATE", "ACTUALSTARTDATE", "ASSIGNEE", "ASSIGNEEGROUP", "CHANGEID", "CLOSURECODE", "COMPLETEDDATE", "ENVIRONMENT", "MODIFIEDDATE", "PRIORITY", "PRODUCTCATEGORIZATIONTIER1", "PRODUCTCATEGORIZATIONTIER2", "PRODUCTCATEGORIZATIONTIER3", "RELATEDCINAME", "RELATEDINCIDENTID", "RFCDATE", "RISKLEVEL", "SCHEDULEDENDDATE", "SCHEDULEDSTARTDATE", "STATUS", "SUBMITDATE", "SUMMARY", "TARGETDATE", "YEAR_MONTH_SCH_DTTM", "HOUR_SCH_DTTM", "DAY_NAME_SCH_DTTM", "YEAR_WEEK_SCH_DTTM", "YEAR_MONTH_ACTST_DTTM", "HOUR_ACTST_DTTM", "DAY_NAME_ACTST_DTTM", "YEAR_MONTH_ACTFT_DTTM", "HOUR_ACTFT_DTTM", "DAY_NAME_ACTFT_DTTM", "YEAR_WEEK_ACTFT_DTTM", "YEAR_MONTH_SCHFT_DTTM", "HOUR_SCHFT_DTTM", "DAY_NAME_SCHFT_DTTM", "COMPARE_START", "COMPARE_FINISH", "AGEBIN", "TTR_DAYS", "TTR_BIN", "FASTPATH_REASON_CD", "ITDCLOSURECODE", "SUCCESSVSFAILED", "ITDEXCEPTION", "PMCHGTYPE", "CANCELLED_COUNT", "CLOSED_COUNT", "SUCCESS_COUNT", "SUCCESSFULPERCENT", "FAIL_COUNT", "FAILEDPERCENT", "EMERGENCY_COUNT", "EMERGENCYPERCENT", "UNAUTH_COUNT", "UNAUTHPERCENT", "APPROVAL_CD", "APPROVED_COUNT", "APPROVEDPERCENT", "REJECTED_COUNT", "REJECTEDPERCENT", "EXCEPTIONPERCENT", "BACKLOG", "COUNT", "YEAR_MONTH_CREATED", "YEAR_MONTH_CLOSED"}
	var chgTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var chgCustName = []string{"Acme", "Acme2"}
	var chgSDataSource = []string{"ACME", "ACME2"}
	var chgRequestDTTM = []string{"06:05.0", "37:31.0", "35:10.0", "52:29.0", "59:22.0", "05:04.0", "44:10.0", "24:20.0", "15:20.0", "24:16.0", "09:38.0", "20:55.0", "43:43.0", "39:43.0", "35:03.0", "22:34.0", "11:15.0", "50:57.0", "23:01.0", "48:30.0", "42:14.0", "55:08.0", "24:17.0", "32:16.0", "04:14.0", "57:02.0", "03:24.0", "49:37.0", "24:41.0", "45:37.0", "34:28.0", "36:22.0", "38:43.0", "20:19.0", "32:02.0", "46:56.0", "59:18.0", "45:00.0", "25:27.0", "39:29.0", "38:55.0", "05:50.0", "18:04.0", "13:00.0", "10:12.0", "30:08.0", "56:10.0", "53:58.0", "22:03.0", "33:01.0", "43:48.0", "45:09.0", "18:51.0", "59:58.0", "42:04.0", "07:00.0", "29:17.0", "41:02.0", "03:56.0", "21:41.0", "13:11.0", "42:19.0", "31:10.0", "28:02.0", "56:09.0", "49:36.0", "59:30.0", "08:31.0", "33:32.0", "20:12.0", "13:08.0", "43:33.0", "56:18.0", "45:34.0", "01:13.0", "15:02.0", "38:56.0", "23:26.0", "45:06.0", "35:13.0", "54:30.0", "45:06.0", "17:54.0", "49:26.0", "31:47.0", "41:36.0", "57:02.0", "01:35.0", "01:36.0", "01:38.0", "01:39.0", "01:41.0", "01:43.0", "07:00.0", "08:33.0", "20:02.0", "20:52.0", "33:04.0", "22:44.0", "33:32.0"}
	var chgPlanStDTTM = []string{"30:00.0", "50:10.0", "30:00.0", "00:00.0", "00:00.0", "30:56.0", "30:00.0", "30:00.0", "30:00.0", "45:00.0", "30:27.0", "00:00.0", "30:00.0", "00:00.0", "", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "30:00.0", "00:00.0", "24:56.0", "00:22.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "", "05:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "30:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "", "30:00.0", "30:00.0", "00:00.0", "00:00.0", "00:00.0", "", "18:36.0", "00:00.0", "00:00.0", "", "00:03.0", "00:00.0", "", "00:00.0", "", "30:00.0", "00:00.0", "15:00.0", "28:59.0", "", "00:00.0", "00:00.0", "47:42.0", "", "00:00.0", "", "00:00.0", "00:00.0", "00:00.0", "41:57.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "00:00.0", "30:00.0", "00:00.0", "", "", "", "00:00.0", ""}
	var chgActStDTTM = []string{"30:01.0", "50:13.0", "30:02.0", "", "00:01.0", "", "30:02.0", "30:04.0", "30:08.0", "45:02.0", "36:25.0", "", "", "00:09.0", "", "00:16.0", "", "00:03.0", "00:12.0", "00:02.0", "00:01.0", "00:01.0", "", "", "", "00:02.0", "00:01.0", "00:01.0", "30:02.0", "00:10.0", "", "00:25.0", "00:03.0", "00:05.0", "00:02.0", "00:01.0", "00:00.0", "", "15:12.0", "00:28.0", "45:03.0", "12:51.0", "17:10.0", "00:12.0", "00:02.0", "", "", "", "30:02.0", "", "00:07.0", "00:05.0", "", "00:08.0", "", "", "", "30:08.0", "00:00.0", "00:03.0", "00:02.0", "", "18:37.0", "00:10.0", "00:03.0", "", "", "00:06.0", "", "00:08.0", "", "30:01.0", "13:22.0", "15:05.0", "31:13.0", "", "00:03.0", "00:10.0", "49:53.0", "", "00:17.0", "", "00:07.0", "00:02.0", "00:04.0", "49:46.0", "00:07.0", "58:01.0", "02:09.0", "02:13.0", "02:21.0", "02:20.0", "02:19.0", "25:56.0", "14:12.0", "", "", "", "00:10.0", ""}
	var chgActFinishDTTM = []string{"19:20.0", "", "13:52.0", "", "46:35.0", "", "59:41.0", "35:57.0", "01:08.0", "22:22.0", "31:00.0", "", "", "44:37.0", "", "18:17.0", "", "52:57.0", "44:24.0", "47:16.0", "08:34.0", "04:40.0", "", "", "", "45:14.0", "43:41.0", "10:42.0", "50:51.0", "50:17.0", "", "", "13:43.0", "25:11.0", "", "30:30.0", "00:00.0", "", "15:23.0", "", "45:56.0", "13:04.0", "17:25.0", "39:32.0", "44:00.0", "", "", "", "", "", "", "38:14.0", "", "41:43.0", "", "", "", "04:37.0", "00:00.0", "14:01.0", "21:46.0", "", "", "", "36:56.0", "", "", "29:25.0", "", "03:25.0", "", "30:00.0", "06:44.0", "57:18.0", "06:34.0", "", "38:53.0", "45:42.0", "36:30.0", "", "47:08.0", "", "44:52.0", "52:29.0", "32:05.0", "13:29.0", "03:36.0", "22:39.0", "23:16.0", "23:57.0", "24:36.0", "26:09.0", "27:11.0", "26:11.0", "49:31.0", "", "", "", "", ""}
	var chgContantSlice = []string{"00:00.0", "00:00.0", "23:34.0", "Applications Software", "Normal", "Closed", "4 - Minor", "VOZ-VA-VAGRY", "yba8rzrnqpi58.rzrn.pbec.wjg.pbz", "", "", "SUCCESSFUL", "17", "1", "", "CHANGE", "0.000277778", "206.508", "Global", "J-COMPANY", "Qvfgevohgrq", "", "", "46:56.0", "fri", "11", "11", "48", "tue", "10", "12", "51", "2", "", "", "", "", "Global", "", "", "", "", "", "24:33.0", "30:30.0", "00:01.0", "x", "VOZ-VA-VAGRY", "CHG5658103", "Unknown", "23:34.0", "", "23:34.0", "3", "", "", "", "", "", "46:56.0", "4", "00:00.0", "00:00.0", "Open", "46:56.0", "", "00:00.0", "2019-12", "3", "sat", "2019-49", "2019-12", "3", "sat", "2019-12", "13", "mon", "2019-51", "2019-12", "23", "sat", "Early in Execution", "Delay in Closing the Change", "", "9.43784", "06-Oct", "", "SUCCESSFUL", "SUCCESSFUL", "NORMAL", "Normal", "0", "1", "1", "100", "0", "0", "0", "0", "0", "0", "APPROVED", "1", "0", "0", "0", "0", "", "1", "2019-11", "2019-12"}

	flag := true
	for gbSize < 200 {
		csvData5chg = nil
		csvData5chg = append(csvData5chg, chgStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand SOURCE_TYPE
			tmpSlice = append(tmpSlice, "DataCenter")

			//rand ID
			rand.Seed(time.Now().UnixNano())
			min := 1000000
			max := 9999999
			randID := rand.Intn(max-min+1) + min
			//fmt.Println(rand.Intn(max-min+1) + min)
			tmpSlice = append(tmpSlice, strconv.Itoa(randID))

			//rand D_CLIENT_ID
			min = 9999123
			max = 9999999
			randClientID := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, strconv.Itoa(randClientID))

			//rand TENANT_ID
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = chgTenantID[rand.Intn(len(chgTenantID))]
			}
			//fmt.Println(TenantIDRune)
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//rand BACID
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand CUSTOMER_NAME
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgCustName[rand.Intn(len(chgCustName))])

			//rand NUMBER
			min = 5000000
			max = 5999999
			randNumber := rand.Intn(max-min+1) + min
			//fmt.Println(randNumber)
			tmpSlice = append(tmpSlice, "CHG"+strconv.Itoa(randNumber))

			//rand S_DATA_SOURCE
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgSDataSource[rand.Intn(len(chgSDataSource))])

			//rand REQUEST_DTTM
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgRequestDTTM[rand.Intn(len(chgRequestDTTM))])

			//rand PLAN_ST_DTTM
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgPlanStDTTM[rand.Intn(len(chgPlanStDTTM))])

			//rand ACT_ST_DTTM
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgActStDTTM[rand.Intn(len(chgActStDTTM))])

			//rand ACT_FINISH_DTTM
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, chgActFinishDTTM[rand.Intn(len(chgActFinishDTTM))])

			for _, row1 := range chgContantSlice {
				tmpSlice = append(tmpSlice, row1)
			}

			//tmpStruct = append(tmpStruct, tmpSlice)

			//create the final slice
			csvData5chg = append(csvData5chg, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData5chg)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData5chg)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/ACME_Chg_data_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData5chg)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "ACME_Chg_data_temp.csv") + " > " + filepath.Join(currentPath, "ACME_Chg_data_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted for change")
		}
		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData5chg)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "ACME_Chg_data_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "change", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-chg.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-chg.json", head, eUser, ePassword, elasticClusterIP, "change")

			//cleanup the file
			cleanup(currentPath, "finalOutput-chg.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-chg.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-chg.json", head, eUser, ePassword, elasticClusterIP, "change")

		//cleanup the file
		cleanup(currentPath, "finalOutput-chg.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/change/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/change/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("Change index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		time.Sleep(2 * time.Second)

		////
		/////
		//////
	}
	defer wg.Done()
}

func dataSetHealth(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData4heal = [][]string{
		{"TENANT_ID", "SOURCE_TYPE", "server", "HEALTH_STATUS", "PLATFORM_TYPE", "provider_account", "correlation_id", "service_category_type"},
	}

	//sample strings health
	var healStartString = []string{"TENANT_ID", "SOURCE_TYPE", "server", "HEALTH_STATUS", "PLATFORM_TYPE", "provider_account", "correlation_id", "service_category_type"}
	var healTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var healServer = []string{"zif7", "zif6", "o7qfnibz9fbnm6.nivfohqtrg.pbz", "o8icjpnvag56.pbec.nivfohqtrg.pbz", "o8iqjvasfcyno56.pbec.nivfohqtrg.pbz", "o8ccjjeyfdy56.pbec.nivfohqtrg.pbz", "o7qfnibz9fbn56.praqnag.pbz", "o8icjvasirnz56.pbec.nivfohqtrg.pbz", "o8cheqonben56.nivfohqtrg.pbz", "o8cjylpcby50.pbec.nivfohqtrg.pbz", "r7ffnotgvzncc57.praqnag.pbz", "o8cfnigecben56.nivfohqtrg.pbz", "o8icjnocbfzta56.pbec.nivfohqtrg.pbz", "o8cfnijngjf57.nivfohqtrg.pbz", "j59icerpzqvep56.jqp59.jqpvoz.nivfohqtrg.pbz", "o78mcejvmncc57.nivfohqtrg.pbz", "o78mcejvmncc56.nivfohqtrg.pbz", "ojudebirejjx8.pbec.nivfohqtrg.pbz", "o7cfpbeuhzvak-g.pbec.nivfohqtrg.pbz", "o8ccecnhben56.nivfohqtrg.pbz", "o8icjvasrkrqt59.pbec.nivfohqtrg.pbz", "o6cfnijroztg56.pbec.nivfohqtrg.pbz", "o8iceqlancc56.nivfohqtrg.pbz", "onopenc568.obhyqre.zrof.vubfg.pbz", "pnopenc054.pbyhzohf.zrof.vubfg.pbz", "pnopejf051.pbyhzohf.zrof.vubfg.pbz", "onopenc567.obhyqre.zrof.vubfg.pbz", "onopenc561.obhyqre.zrof.vubfg.pbz", "pnopeqo059.pbyhzohf.zrof.vubfg.pbz", "onopeqo059.obhyqre.zrof.vubfg.pbz", "pnopenc066.pbyhzohf.zrof.vubfg.pbz", "pnopeqo550.pbyhzohf.zrof.vubfg.pbz", "pnopenc568.pbyhzohf.zrof.vubfg.pbz", "o8cheqyjben56.nivfohqtrg.pbz", "o8icjnoorffdy57.pbec.nivfohqtrg.pbz", "o8ccrbztfs57.pbec.nivfohqtrg.pbz", "o8ccjvasrkz57.pbec.nivfohqtrg.pbz", "o8ccjvasrkz56.pbec.nivfohqtrg.pbz", "o7cfongtqftgj51.pbec.nivfohqtrg.pbz", "o8icjvasrkqnt56.pbec.nivfohqtrg.pbz", "o8ihjylpncc57.grfgpbec.nivfohqtrg.grfg", "o8cceqyancc56.nivfohqtrg.pbz", "ovvcfqupc59.pbec.nivfohqtrg.pbz", "ovcvasgqztgjl56.pbec.nivfohqtrg.pbz", "o8cfnjmyqncc58.nivfohqtrg.pbz", "o8cfnjmyqncc57.nivfohqtrg.pbz", "r7fqnipfuejgf57.pbec.nivfohqtrg.pbz", "o8cceqclncc56.nivfohqtrg.pbz", "o8ccrfsrfk50.pbec.nivfohqtrg.pbz", "awcfvasoxcgfz56.praqnag.pbz", "oc6cfvasrfk57.nivfohqtrg.pbz", "o8icjpgkkra58.pbec.nivfohqtrg.pbz", "o8icjpgkkra50.pbec.nivfohqtrg.pbz", "o8chrbztfs57.pbec.nivfohqtrg.pbz", "o8chrbztfs56.pbec.nivfohqtrg.pbz", "o6icynopucncc56.pbec.nivfohqtrg.pbz", "o8cfjjeyfdy50.pbec.nivfohqtrg.pbz", "onopenc554.obhyqre.zrof.vubfg.pbz", "o8ccjppnencc56.pbec.nivfohqtrg.pbz", "o8ccjppnencc57.pbec.nivfohqtrg.pbz", "o8cherzrncc56.nivfohqtrg.pbz", "q65igjrpzwzc56.pbec.nivfohqtrg.pbz", "o8cceqscben56.nivfohqtrg.pbz", "o8cheqscncc57.nivfohqtrg.pbz", "o8icjnotbbqqo65.pbec.nivfohqtrg.pbz", "o8cceqyjben56.nivfohqtrg.pbz", "awcjpyhfgsf56.pbec.nivfohqtrg.pbz", "o6icynopucncc57.pbec.nivfohqtrg.pbz", "o8ccjuczcfchy57.pbec.nivfohqtrg.pbz", "q65icefueegz56.qny65.qnyvoz.nivfohqtrg.pbz", "q65icerpzncc57.qny65.qnyvoz.nivfohqtrg.pbz", "o8cckvasgfz59.nivfohqtrg.pbz", "q65igerpzncc58.qny65.qnyvoz.nivfohqtrg.pbz", "o8ccrbztfu57.pbec.nivfohqtrg.pbz", "ugqp57-221qp154.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icerpznrzn56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icerpznrzn57.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icefueqbc56.jqp59.jqpvoz.nivfohqtrg.pbz", "o6cfninccecg56.pbec.nivfohqtrg.pbz", "j59icefueegz57.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icjfuenq56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz52.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzncc56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icjrpzwzc56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz58.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzrqo58.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzjro57.jqp59.jqpvoz.nivfohqtrg.pbz", "o8ccrfsrfk59.pbec.nivfohqtrg.pbz", "o8ccjfnifsosr57.pbec.nivfohqtrg.pbz", "o8icjpnebbg56.pbec.nivfohqtrg.pbz", "o8ccecrzben56.nivfohqtrg.pbz", "q65igerpzncc57.qny65.qnyvoz.nivfohqtrg.pbz", "q65igerpzrqo58.qny65.qnyvoz.nivfohqtrg.pbz", "q65igerpzncc56.qny65.qnyvoz.nivfohqtrg.pbz", "ugqp56-221qp154.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icerpzrqo58.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icerpzncc56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icerpznrzc57.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icefueybt56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz67.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzncc58.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icefuezckl56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59icjrpzngg56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz53.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzncc59.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igefuerqo56.jqp59.jqpvoz.nivfohqtrg.pbz", "pqvasen6.nivfohqtrg.pbz", "o8ccrfsrfk56.pbec.nivfohqtrg.pbz", "o8ihebztyqnc56.nivfohqtrg.pbz", "o8icjavprnc56.pbec.nivfohqtrg.pbz", "o7cfnibz0fbn57.praqnag.pbz", "o8iqbztgbby56.nivfohqtrg.pbz", "otmqfnipfuefei56.nivfohqtrg.pbz", "o8icjecgpzncc56.pbec.nivfohqtrg.pbz", "q65icjfrpwzc56.pbec.nivfohqtrg.pbz", "r7cfnotvfnckl56.pbec.nivfohqtrg.pbz", "pnopenc564.pbyhzohf.zrof.vubfg.pbz", "pnopenc554.pbyhzohf.zrof.vubfg.pbz", "o7cfnotbz0qo58.nivfohqtrg.pbz", "o8ccrfsrfk51.pbec.nivfohqtrg.pbz", "o8icebztyqnc57.nivfohqtrg.pbz", "ovcfvasavz96.nivfohqtrg.pbz", "o8cheqyaben56.nivfohqtrg.pbz", "o8chrzqzrfk56", "onopenc569.obhyqre.zrof.vubfg.pbz", "o8cqenebflo66.nivfohqtrg.pbz", "o78mcejvmzba57.nivfohqtrg.pbz", "o8chjqclfdy56.pbec.nivfohqtrg.pbz", "o8icjvasacf56.pbec.nivfohqtrg.pbz", "o8icjvasacf57.pbec.nivfohqtrg.pbz", "o8cheqfgncc56.nivfohqtrg.pbz", "o8cheqfgncc57.nivfohqtrg.pbz", "o8cheqypncc57.nivfohqtrg.pbz", "rqtr7.nivfohqtrg.pbz", "o8cceqypben56.nivfohqtrg.pbz", "o8cceqfgben56.nivfohqtrg.pbz", "o8cceqclncc57.nivfohqtrg.pbz", "o8ihjhvcjro57.pbec.nivfohqtrg.pbz", "qo-c6.nivfohqtrg.pbz", "o8cceqclben56.nivfohqtrg.pbz", "awccrrfk56.pbec.nivfohqtrg.pbz", "ojud-ebire-jjx6.pbec.nivfohqtrg.pbz", "o8cceqonben56.nivfohqtrg.pbz", "o8cheqonben57.nivfohqtrg.pbz", "o8ccenqoben56.nivfohqtrg.pbz", "o8icjjeyjro66.pbec.nivfohqtrg.pbz", "o8ffprqjc-qo6.nivfohqtrg.pbz", "ofmqfnipfueqri59.nivfohqtrg.pbz", "o8ccjfnifsorqt6.pbec.nivfohqtrg.pbz", "o8ccjfnifsorqt7.pbec.nivfohqtrg.pbz", "otmcfnipjrofei56.nivfohqtrg.pbz", "o7qfnibz9ecg56.pbec.nivfohqtrg.pbz", "ipragre-jqp59.jqp59.jqpvoz.nivfohqtrg.pbz", "o8icewhzcubfg56.nivfohqtrg.pbz", "o8icepngqon56.nivfohqtrg.pbz", "o8icepngqon57.nivfohqtrg.pbz", "r7cfpbejrockl57.nivfohqtrg.pbz", "o8icebztncc56.nivfohqtrg.pbz", "o8icebztncc57.nivfohqtrg.pbz", "o8icebztncc59.nivfohqtrg.pbz", "o8ihebztncc56.nivfohqtrg.pbz", "o8ihebztncc57.nivfohqtrg.pbz", "o8icebztncc53.nivfohqtrg.pbz", "o8cqeqscazd57.nivfohqtrg.pbz", "o8checebncc69.nivfohqtrg.pbz", "awccrrfk59.pbec.nivfohqtrg.pbz", "o7qfni9fbntom.nivfohqtrg.pbz", "inosf58.pbec.nivfohqtrg.pbz", "o8icjnotpgkkn51.pbec.nivfohqtrg.pbz", "o8qfnigvzncc56.nivfohqtrg.pbz", "j59icerpzpnu56.jqp59.jqpvoz.nivfohqtrg.pbz", "o8checebben68.nivfohqtrg.pbz", "o8chepzfben56.nivfohqtrg.pbz", "o8ccenebncc68.nivfohqtrg.pbz", "o8iqjdnncc58.pbec.nivfohqtrg.pbz", "o8icjnepuefei56.pbec.nivfohqtrg.pbz", "o8chepzfncc56.nivfohqtrg.pbz", "o8cheqyjncc56.nivfohqtrg.pbz", "o8cheqyjncc57.nivfohqtrg.pbz", "q65icjfuewzc56.pbec.nivfohqtrg.pbz", "o8ccecebben66.nivfohqtrg.pbz", "o8ccrfsrfk52.nivfohqtrg.pbz", "awicjpherncc56.pbec.nivfohqtrg.pbz", "o8icjbzincc58.pbec.nivfohqtrg.pbz", "pnopenc060.pbyhzohf.zrof.vubfg.pbz", "o8ccrbztfs56.pbec.nivfohqtrg.pbz", "o8icebztncc51.nivfohqtrg.pbz", "nmcfnipbecqp57.pbec.nivfohqtrg.pbz", "awccrrfk57.pbec.nivfohqtrg.pbz", "o8icjnotpgkkn52.pbec.nivfohqtrg.pbz", "o8icjpgkkra53.pbec.nivfohqtrg.pbz", "o8ccegjben57.nivfohqtrg.pbz", "o8ccrwhzcfu57.pbec.nivfohqtrg.pbz", "o8hfnic3pc56.nivfohqtrg.pbz", "q65igerpznrz52.qny65.qnyvoz.nivfohqtrg.pbz", "o7cfpbeuhzncc56.pbec.nivfohqtrg.pbz", "o7cfpbeuhzfdy56.pbec.nivfohqtrg.pbz", "j59icefueegz56.jqp59.jqpvoz.nivfohqtrg.pbz", "o8icjjeyjro56.pbec.nivfohqtrg.pbz", "pnopeqo551.pbyhzohf.zrof.vubfg.pbz", "onopenc564.obhyqre.zrof.vubfg.pbz", "o7cfnotbz0qo57.nivfohqtrg.pbz", "o8icjorf67ncc56.pbec.nivfohqtrg.pbz", "o6cnierpbafdy56.pbec.nivfohqtrg.pbz", "o8ccjppnejro56.pbec.nivfohqtrg.pbz", "o7cfnryzyttgj57.pbec.nivfohqtrg.pbz", "o6cfnincc8ie58.pbec.nivfohqtrg.pbz", "o8icepngqon58.nivfohqtrg.pbz", "o8cfnibz0df56.nivfohqtrg.pbz", "o8cfjjeyfdy58.pbec.nivfohqtrg.pbz", "r7cfnotyzptgj57.praqnag.pbz", "o7cfnipzyttgj57.pbec.nivfohqtrg.pbz", "pnopenc053.pbyhzohf.zrof.vubfg.pbz", "o8iqeqonben56.nivfohqtrg.pbz", "j59igjrpzwzc56.pbec.nivfohqtrg.pbz", "o8cqrbztfu56.pbec.nivfohqtrg.pbz", "o8icjnocbfzta58.pbec.nivfohqtrg.pbz", "o8ccjcnhfdy56.pbec.nivfohqtrg.pbz", "o6cfnopragjro56.pbecqzm.ypy", "o7cfnipzyttgj53.pbec.nivfohqtrg.pbz", "pnopenc553.pbyhzohf.zrof.vubfg.pbz", "j59icefueybt57.jqp59.jqpvoz.nivfohqtrg.pbz", "o8cjylpcby59.pbec.nivfohqtrg.pbz", "o7cfnotgjvncc56.nivfohqtrg.pbz", "q65icjrpzwzc56.pbec.nivfohqtrg.pbz", "q65icjfuenq57.pbec.nivfohqtrg.pbz", "o8icjpoffdy56.pbec.nivfohqtrg.pbz", "o8iqjfceg6857.pbec.nivfohqtrg.pbz", "j59icerpzpnu57.jqp59.jqpvoz.nivfohqtrg.pbz", "r7qfnotbz8ncc56.praqnag.pbz", "oc6cfvasrfk56.nivfohqtrg.pbz", "o7cfohqzyttgj57.pbec.nivfohqtrg.pbz", "o7cffnozyttgj57.pbec.nivfohqtrg.pbz", "g6cfohqncpqo56.jvmpbz.pbec.nivfohqtrg.pbz", "o7cfnryzyttgj56.pbec.nivfohqtrg.pbz", "ovcfvasfucncc56.pbec.nivfohqtrg.pbz", "r7cfnotlyqncc56.nivfohqtrg.pbz", "or7qfnisri8frp58.nivfohqtrg.pbz", "o7qfnipjbjncc56.nivfohqtrg.pbz", "r7cfnotbz0ncc56.praqnag.pbz", "o7cfnzqzyttgj57.pbec.nivfohqtrg.pbz", "ovqzmvfn57.pbec.nivfohqtrg.pbz", "o7cfnipjylncc56.nivfohqtrg.pbz", "ofmdfnipfuezfp56.nivfohqtrg.pbz", "ofmdfnipfuezfp57.nivfohqtrg.pbz", "o6cfnipragyqf56.pbec.nivfohqtrg.pbz", "ofmqfniprqjqri56.nivfohqtrg.pbz", "g6cfprapoffdy56.pbec.nivfohqtrg.pbz", "o6cfnifrcva56.pbec.nivfohqtrg.pbz", "o7qfnotbz9qo56.nivfohqtrg.pbz", "o7qfnotbz9qo57.nivfohqtrg.pbz", "o8cfbofcragrfgiz56", "r7cfnipoymncc67.pbec.nivfohqtrg.pbz", "o6qfnisfpncc56.pbec.nivfohqtrg.pbz", "oemcfniplyqqo56.nivfohqtrg.pbz", "o7cfnohngjro56.pbec.nivfohqtrg.pbz", "o8qfnotvas56.nivfohqtrg.pbz", "o8cfnotvas56.nivfohqtrg.pbz", "o8tmcfniperi57.nivfohqtrg.pbz", "o8hfnic3pyqz56.pbec.nivfohqtrg.pbz", "oemcfnipfuezfp56.nivfohqtrg.pbz", "r7cfnotbz0kzy57.nivfohqtrg.pbz", "r7cfnotbz0kzy56.nivfohqtrg.pbz", "o6cfjvmfuevdqo56.pbec.nivfohqtrg.pbz", "o6cfnicebkgnq56.pbec.nivfohqtrg.pbz", "o6cfniqoffdy57.pbec.nivfohqtrg.pbz", "ofmqfnipfueqri56.nivfohqtrg.pbz", "o6cfniieffdy57.pbec.nivfohqtrg.pbz", "ofmjepqri57.nivfohqtrg.pbz", "oqcvasgqztgjl56.pbecnoqzm.ypy", "o8qfnipzoncc56.pbec.nivfohqtrg.pbz", "o8hfniic3jro56.pbec.nivfohqtrg.pbz", "o8cfniic3jro56.pbec.nivfohqtrg.pbz", "o8cfniic3jro57.pbec.nivfohqtrg.pbz", "o8checnhben56.nivfohqtrg.pbz", "onopenc065.obhyqre.zrof.vubfg.pbz", "o8qfnotqrincc56.nivfohqtrg.pbz", "ov6cfvasipqo56.pbec.nivfohqtrg.pbz", "o8cfnotfgtqo56.nivfohqtrg.pbz", "o8hfnic3nr57.nivfohqtrg.pbz", "o8cfnic3nr56.nivfohqtrg.pbz", "o8qfniegpben56.nivfohqtrg.pbz", "jud-jvmpbz-oqp6.pbec.nivfohqtrg.pbz", "o7cfnippngncc56.pbec.nivfohqtrg.pbz", "ovcfvasqqp57.pbec.nivfohqtrg.pbz", "ovcfvasegqp57.pbec.nivfohqtrg.pbz", "o8icjpgkyvp56.pbec.nivfohqtrg.pbz", "ovcftevaspgk56.pbec.nivfohqtrg.pbz", "ovcftevaspgk57.pbec.nivfohqtrg.pbz", "o78mhejvmncc56.nivfohqtrg.pbz", "r7cfnotbz0vag57.praqnag.pbz", "o8icjcnlyffdy56.pbec.nivfohqtrg.pbz", "o78mcejvmzba56.nivfohqtrg.pbz", "o8cceqscncc57.nivfohqtrg.pbz", "o8icjecgjfncc58.pbec.nivfohqtrg.pbz", "o8cceqyaben56.nivfohqtrg.pbz", "o7cfvaseqfvnf56.pbec.nivfohqtrg.pbz", "r7cfnipsygjgf57.pbec.nivfohqtrg.pbz", "o8iqjohffdy56.pbec.nivfohqtrg.pbz", "o8ccecebben68.nivfohqtrg.pbz", "o8icjjeyjro57.pbec.nivfohqtrg.pbz", "o8cfjjeyfdy59.pbec.nivfohqtrg.pbz", "o8cceenhncc56.nivfohqtrg.pbz", "o8cfjjeyfdy57.pbec.nivfohqtrg.pbz", "o8iqjdnqo56.pbec.nivfohqtrg.pbz", "o8ccjjeyfdy59.pbec.nivfohqtrg.pbz", "o8chjjeyfdy56.pbec.nivfohqtrg.pbz", "ofmcfnipfuejro56.nivfohqtrg.pbz", "j59icjfuenq57.pbec.nivfohqtrg.pbz", "o8ccenqoben57.nivfohqtrg.pbz", "o8cfnibz0qgpby6.nivfohqtrg.pbz", "o8icjpgkkra59.pbec.nivfohqtrg.pbz", "o8checebben67.nivfohqtrg.pbz", "o8chrbztfu56.pbec.nivfohqtrg.pbz", "o8ihjhvcfdy56.pbec.nivfohqtrg.pbz", "j59icerpznrzc56.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpzncc57.jqp59.jqpvoz.nivfohqtrg.pbz", "pnopenc069.pbyhzohf.zrof.vubfg.pbz", "o8icjpgkkra57.pbec.nivfohqtrg.pbz", "awcjpyhfgsf57.pbec.nivfohqtrg.pbz", "o8ccjppnejro57.pbec.nivfohqtrg.pbz", "j59icjfuewzc56.pbec.nivfohqtrg.pbz", "o7cfnohngfdy56.pbec.nivfohqtrg.pbz", "o8icebztfvzncc57.nivfohqtrg.pbz", "onopenc560.obhyqre.zrof.vubfg.pbz", "j59igerpznrz65.jqp59.jqpvoz.nivfohqtrg.pbz", "o8checebben66.nivfohqtrg.pbz", "o7qfnotbz8qo56.nivfohqtrg.pbz", "oemcfnipfueqon56.nivfohqtrg.pbz", "o8ccecebncc66.nivfohqtrg.pbz", "r7qfnotbz9vag57.praqnag.pbz", "o8ccjfnifsoqve6.pbec.nivfohqtrg.pbz", "o8cfnjvmjf59.nivfohqtrg.pbz", "onopenc553.obhyqre.zrof.vubfg.pbz", "onopeqo551.obhyqre.zrof.vubfg.pbz", "o7qfnibz9fdy56.pbec.nivfohqtrg.pbz", "o8ccjfnifsopdq6.pbec.nivfohqtrg.pbz", "r7qfnotbz9ncc56.praqnag.pbz", "o8cceqscncc56.nivfohqtrg.pbz", "o6cfninccpbe56.pbec.nivfohqtrg.pbz", "ocvcfvasfwf56.praqnag.pbz", "r7qfnotyzptgj56.praqnag.pbz", "r7cfnipoymqri59.pbec.nivfohqtrg.pbz", "o6cfniibygfdy56.pbec.nivfohqtrg.pbz", "o7cfnibz0fbn56.praqnag.pbz", "o8ccerzrncc57.nivfohqtrg.pbz", "o8icjecgjfncc59.pbec.nivfohqtrg.pbz", "o8hfnifzfncc56.nivfohqtrg.pbz", "o8icjpnvff56.pbec.nivfohqtrg.pbz", "r7cfnipfvtpnc56.nivfohqtrg.pbz", "o8icjjeyjro54.pbec.nivfohqtrg.pbz", "ovcfvasfucfdy56.pbec.nivfohqtrg.pbz", "o8ccecebncc60.nivfohqtrg.pbz", "o8icjnotpgkkn59.pbec.nivfohqtrg.pbz", "o8ccjgcpfpp56.pbec.nivfohqtrg.pbz", "pnopeqo050.pbyhzohf.zrof.vubfg.pbz", "pqvasen7.nivfohqtrg.pbz", "o6cfnincc8ie57.pbec.nivfohqtrg.pbz", "o8chenqoben56.nivfohqtrg.pbz", "o8ccepzfben56.nivfohqtrg.pbz", "o7cfjfczyttgj56.pbec.nivfohqtrg.pbz", "o8cfnibz0zd57.nivfohqtrg.pbz", "o8qfnotj8gf57.nivfohqtrg.pbz", "o8cfnipxnapfr56.nivfohqtrg.pbz", "o8ccrfsrfk53.nivfohqtrg.pbz", "rfkv7.qny65.qnyvoz.nivfohqtrg.pbz", "o8cceqfoben56.nivfohqtrg.pbz", "o8cheqclben56.nivfohqtrg.pbz", "o8cjvasfxyz59.pbec.nivfohqtrg.pbz", "pnopeqo559.pbyhzohf.zrof.vubfg.pbz", "r7cfnotbz0ncc57.praqnag.pbz", "o8ccjvasrkz58.pbec.nivfohqtrg.pbz", "o8cheenhncc56.nivfohqtrg.pbz", "o8qfnibz9zqz56.nivfohqtrg.pbz", "o8icjecgjfncc56.pbec.nivfohqtrg.pbz", "o8cfnic3pc56.nivfohqtrg.pbz", "o8icjecgpzncc57.pbec.nivfohqtrg.pbz", "o8icjecgjfncc57.pbec.nivfohqtrg.pbz", "o8icjvasjro57.pbec.nivfohqtrg.pbz", "j59icerpzrqo57.jqp59.jqpvoz.nivfohqtrg.pbz", "o8ccrwhzcfu56.pbec.nivfohqrg.pbz", "o8ccjvasrkz51.pbec.nivfohqtrg.pbz", "o8icyvasip56.pbec.nivfohqtrg.pbz", "q65icjfuefcir56.pbec.nivfohqtrg.pbz", "o7cfnotouq56.nivfohqtrg.pbz", "o8icjnoorffdy56.pbec.nivfohqtrg.pbz", "g6cjavprherp56.pbec.nivfohqtrg.pbz", "o8ccexnapfr50.nivfohqtrg.pbz", "o8qfnotj8gf56.nivfohqtrg.pbz", "o8cheqscben56.nivfohqtrg.pbz", "o8icjpgkkra54.pbec.nivfohqtrg.pbz", "o8qfnipynben56.nivfohqtrg.pbz", "q65icerpzjns57.qny65.qnyvoz.nivfohqtrg.pbz", "o8ccrecgrfk57.pbec.nivfohqtrg.pbz", "o8cjylppzf58.pbec.nivfohqtrg.pbz", "o8icjavprratr56.pbec.nivfohqtrg.pbz", "o8cceqonben58.nivfohqtrg.pbz", "r7qfnotbz9vag56.praqnag.pbz", "o7cfnotzaqncc56.pbec.nivfohqtrg.pbz", "j59igerpzrqo57.jqp59.jqpvoz.nivfohqtrg.pbz", "o7qfnotzaqncc56.pbec.nivfohqtrg.pbz", "o8iqebztncc56.nivfohqtrg.pbz", "o8icjhvcjro57.pbec.nivfohqtrg.pbz", "ovcfvasfucjro56.pbec.nivfohqtrg.pbz", "o8cjvasfxyz58.pbec.nivfohqtrg.pbz", "afkznantre.qny65.qnyvoz.nivfohqtrg.pbz", "q65igerpznrz56.qny65.qnyvoz.nivfohqtrg.pbz", "onopeqo050.obhyqre.zrof.vubfg.pbz", "o8qjfcecg56.pbec.nivfohqtrg.pbz", "o8cfnibz0zqz57.nivfohqtrg.pbz", "o8ffprqjc-qo7.nivfohqtrg.pbz", "pnopejf050.pbyhzohf.zrof.vubfg.pbz", "j59icjrpzwzc56.pbec.nivfohqtrg.pbz", "q65icerpznrzc57.qny65.qnyvoz.nivfohqtrg.pbz", "oemcfninebfym56.nivfohqtrg.pbz", "pnopenc575.pbyhzohf.zrof.vubfg.pbz", "o8ccrfsrfk57.nivf.ohqtrg.pbz", "otmdfnipfuefei56.nivfohqtrg.pbz", "o8icebztncc58.nivfohqtrg.pbz", "r7cfnipfvtpnc57.nivfohqtrg.pbz", "j59icjfrpwzc56.pbec.nivfohqtrg.pbz", "q65icerpzncc56.qny65.qnyvoz.nivfohqtrg.pbz", "q65igerpzjns56.qny65.qnyvoz.nivfohqtrg.pbz", "q65icerpznrzc59.qny65.qnyvoz.nivfohqtrg.pbz", "j59icjfuefcir56.jqp59.jqpvoz.nivfohqtrg.pbz", "o8cjylppzf51.pbec.nivfohqtrg.pbz", "pnopenc567.pbyhzohf.zrof.vubfg.pbz", "onopenc053.obhyqre.zrof.vubfg.pbz", "q65iceoxcgfz56.qny65.qnyvoz.nivfohqtrg.pbz", "q65icerpzqvep56.qny65.qnyvoz.nivfohqtrg.pbz", "o8ccecebben67.nivfohqtrg.pbz", "o8ccjppnencc58.pbec.nivfohqtrg.pbz", "o8cfnic3pc57.nivfohqtrg.pbz", "pnopenc068.pbyhzohf.zrof.vubfg.pbz", "o8hfnic3nr56.nivfohqtrg.pbz", "q65icerpzpnu56.qny65.qnyvoz.nivfohqtrg.pbz", "ojud-ebire-jjx7.pbec.nivfohqtrg.pbz", "o7cfpbeuhzncc-g.pbec.nivfohqtrg.pbz", "ovqppraqnagqzm7.pbecnoqzm.ypy", "j59igerpznrz50.jqp59.jqpvoz.nivfohqtrg.pbz", "r7cfnotbz0vag56.praqnag.pbz", "o8chjjeyfdy57.pbec.nivfohqtrg.pbz", "o8ccjfnifsopzf6.pbec.nivfohqtrg.pbz", "o8icebztfvzncc56.nivfohqtrg.pbz", "o8icjjeqfyf56.pbec.nivfohqtrg.pbz", "j59icefuenczqo56.pbec.nivfohqtrg.pbz", "o7cfpbeohfncc56.pbec.nivfohqtrg.pbz", "onopenc575.obhyqre.zrof.vubfg.pbz", "o8cheqscncc56.nivfohqtrg.pbz", "o8ihjjeyncc56.pbec.nivfohqtrg.pbz", "nmcfnipbecqp56.pbec.nivfohqtrg.pbz", "o8ccjjeyfdy58.pbec.nivfohqtrg.pbz", "awccrrfk58.pbec.nivfohqtrg.pbz", "j59iceoxcgfz56.jqp59.jqpvoz.nivfohqtrg.pbz", "o8ccecnhncc56.nivfohqtrg.pbz", "onopeqo550.obhyqre.zrof.vubfg.pbz", "o7cfpbeohfncc58.pbec.nivfohqtrg.pbz", "r7qfnotbz9ncc57.praqnag.pbz", "onopenc061.obhyqre.zrof.vubfg.pbz", "o7cfpbeuhzfdy-g.nivfohqtrg.pbz", "o8ccegjben56.nivfohqtrg.pbz", "o7cfnotbz0qo56.nivfohqtrg.pbz", "o8ccrbztfu59.pbec.nivfohqtrg.pbz", "rfkv0.jqp59.jqpvoz.nivfohqtrg.pbz", "o8cceqonben57.nivfohqtrg.pbz", "rfkv8.jqp59.jqpvoz.nivfohqtrg.pbz", "rfkv9.qny65.qnyvoz.nivfohqtrg.pbz", "o8ccecebncc68.nivfohqtrg.pbz", "o8icebztncc50.nivfohqtrg.pbz", "g6ccevasgfz56.nivfohqtrg.pbz", "o8icjnocbfzta57.pbec.nivfohqtrg.pbz", "o8tmcfniperi56.nivfohqtrg.pbz", "o8icjpgkqo56.pbec.nivfohqtrg.pbz", "o8ccecebncc69.nivfohqtrg.pbz", "rfkv5.jqp59.jqpvoz.nivfohqtrg.pbz", "j59igerpznrz51.jqp59.jqpvoz.nivfohqtrg.pbz", "cc-ermcnex.nivfohqtrg.pbz", "o8ihjhvcjro56.pbec.nivfohqtrg.pbz", "o8icjgrz56.pbec.nivfohqtrg.pbz", "o8checebncc68.nivfohqtrg.pbz", "q65icerpznrzc56.pbec.nivfohqtrg.pbz", "q65icerpzrqo58.qny65.qnyvoz.nivfohqtrg.pbz", "o8icjvasrkrqt58.pbec.nivfohqtrg.pbz", "o6cfniieffdy56.pbec.nivfohqtrg.pbz", "o8cfnifzfncc56.nivfohqtrg.pbz", "pnopenc576.pbyhzohf.zrof.vubfg.pbz", "o8ccrfsrfk58.pbec.nivfohqtrg.pbz", "o8ccegjvncc56.nivfohqtrg.pbz", "rfkv1.jqp59.jqpvoz.nivfohqtrg.pbz", "o7qfnohngztg56.pbec.nivfohqtrg.pbz", "o8icjsbqscy56.pbec.nivfohqtrg.pbz", "rfkv6.qny65.qnyvoz.nivfohqtrg.pbz", "r7cfpbejrockl56.nivfohqtrg.pbz", "o7cfpbeuhzvak56.pbec.nivfohqtrg.pbz", "o8checnhncc56.nivfohqtrg.pbz", "o7cfjfczyttgj57.pbec.nivfohqtrg.pbz", "o8ccenavncc57.nivfohqtrg.pbz", "o8cfniylprqt56.pbecqzm.ypy", "j59igerpzrqo50.jqp59.jqpvoz.nivfohqtrg.pbz", "o8cfnipzffdy56.pbec.nivfohqtrg.pbz", "o8iqjohffdy-g.pbec.nivfohqtrg.pbz"}
	var healHealthStatus = []string{"Healthy", "Critical", "Warning"}
	var healPlatformType = []string{"Mainframe", "Server"}

	flag := true
	for gbSize < 200 {
		csvData4heal = nil
		csvData4heal = append(csvData4heal, healStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand TENANT_ID
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = healTenantID[rand.Intn(len(healTenantID))]
			}
			//fmt.Println(TenantIDRune)
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//rand SOURCE_TYPE
			tmpSlice = append(tmpSlice, "DataCenter")

			//rand server
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, healServer[rand.Intn(len(healServer))])

			//rand HEALTH_STATUS
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, healHealthStatus[rand.Intn(len(healHealthStatus))])

			//rand PLATFORM_TYPE
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			platformType := healPlatformType[rand.Intn(len(healPlatformType))]
			tmpSlice = append(tmpSlice, platformType)

			//rand provider_account
			min := 1000000
			max := 9999999
			randNumber := rand.Intn(max-min+1) + min
			//fmt.Println(randNumber)
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randNumber))

			//rand correlation_id
			tmpSlice = append(tmpSlice, "IBM-"+string(TenantIDRune)+"-100000000-069152")

			//rand service_category_type
			tmpSlice = append(tmpSlice, "DataCenter::"+platformType+"::")

			//create the final slice
			csvData4heal = append(csvData4heal, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData4heal)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData4heal)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/ACME_hybridhealthhistory_data_all_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData4heal)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "ACME_hybridhealthhistory_data_all_temp.csv") + " > " + filepath.Join(currentPath, "ACME_hybridhealthhistory_data_all_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted for health")
		}
		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData4heal)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "ACME_hybridhealthhistory_data_all_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "health", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-health.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-health.json", head, eUser, ePassword, elasticClusterIP, "health")

			//cleanup the file
			cleanup(currentPath, "finalOutput-health.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-health.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-health.json", head, eUser, ePassword, elasticClusterIP, "health")

		//cleanup the file
		cleanup(currentPath, "finalOutput-health.json")

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/health/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			flag = false
		}

		/////////

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/health/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("Heath index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		time.Sleep(2 * time.Second)

		////
		/////
		//////
	}
	defer wg.Done()
}

func dataSetProblem(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {
	var csvData3prb = [][]string{
		{"ABSTRACT", "ACT_FINISH_DTTM", "ACT_ST_DTTM", "AGEBIN", "BACID", "BACKLOG", "BACKLOGPERCENT", "BIN", "CALL_CD", "CANCELLED_DF", "CATEGORY", "CAUSE", "CI", "CLOSEDTICKET", "CLOSED_DATE_DF", "CLOSED_DF", "CLOSED_DTTM", "CLOSE_CODE", "COUNT", "CREATIONDATE_DF", "DAY_NAME_CREATED", "DAY_NAME_RESOLVED", "DV_SEVERITY", "D_CLIENT_ID", "EFFICIENCYPERCENT", "FINAL_STATUS", "FOLDER_PATH", "HOUR_CREATED", "HOUR_RESOLVED", "ID", "INCIDENT_CODE_ID", "LABEL", "LAST_REFRESH_DTTM", "MODIFY_DTTM", "NO_OWNER_DF", "NO_OWNER_PR", "OCCURRED_DTTM", "OPENED_DATE_DF", "OPEN_DTTM", "ORIG_RECORD", "ORIG_RECORD_CLASS", "OWNER", "QUEUE_ID", "RCA_COMPLETE", "REASSIGNMENTS", "REOPEN_CUSTOM", "RESOLUTION_CODE", "RESOLVED_DTTM", "RESOLVER_NAME", "SEVERITY", "SL", "SLA_STATUS", "SOURCE_TYPE", "STATUS", "STATUS_DF", "SUBCOMPONENT", "S_DATA_SOURCE", "TENANT_ID", "TTR_BIN", "TTR_DAYS", "YEAR_MONTH_CREATED", "YEAR_MONTH_RESOLVED", "YEAR_WEEK_CREATED"},
	}

	//sample strings problem
	var prbStartString = []string{"ABSTRACT", "ACT_FINISH_DTTM", "ACT_ST_DTTM", "AGEBIN", "BACID", "BACKLOG", "BACKLOGPERCENT", "BIN", "CALL_CD", "CANCELLED_DF", "CATEGORY", "CAUSE", "CI", "CLOSEDTICKET", "CLOSED_DATE_DF", "CLOSED_DF", "CLOSED_DTTM", "CLOSE_CODE", "COUNT", "CREATIONDATE_DF", "DAY_NAME_CREATED", "DAY_NAME_RESOLVED", "DV_SEVERITY", "D_CLIENT_ID", "EFFICIENCYPERCENT", "FINAL_STATUS", "FOLDER_PATH", "HOUR_CREATED", "HOUR_RESOLVED", "ID", "INCIDENT_CODE_ID", "LABEL", "LAST_REFRESH_DTTM", "MODIFY_DTTM", "NO_OWNER_DF", "NO_OWNER_PR", "OCCURRED_DTTM", "OPENED_DATE_DF", "OPEN_DTTM", "ORIG_RECORD", "ORIG_RECORD_CLASS", "OWNER", "QUEUE_ID", "RCA_COMPLETE", "REASSIGNMENTS", "REOPEN_CUSTOM", "RESOLUTION_CODE", "RESOLVED_DTTM", "RESOLVER_NAME", "SEVERITY", "SL", "SLA_STATUS", "SOURCE_TYPE", "STATUS", "STATUS_DF", "SUBCOMPONENT", "S_DATA_SOURCE", "TENANT_ID", "TTR_BIN", "TTR_DAYS", "YEAR_MONTH_CREATED", "YEAR_MONTH_RESOLVED", "YEAR_WEEK_CREATED"}
	var prbTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var prbCallCD = []string{"PHONECALL", "PHONE"}
	var prbCategory = []string{"UNKNOWN", "DEVELOPMENT", "PRE_PRODUCTION", "DEVELOPMENT", "PRODUCTION", "DEVELOPMENT", "STAGING", "UNKNOWN", "TEST"}
	var prbCause = []string{"HARDWARE", "SOFTWARE", ""}
	var prbCloseCode = []string{"null", ""}
	var prbOriRecClass = []string{"INCIDENT", "CHANGE"}
	var prbQueueID = []string{"HAXABJA", "NOPE-V-ARGFCQRA", "NOPE-V-IZJNER", "NOPE-V-SVERJNYY", "NOPE-V-VAQJVAGRY", "NOPE-V-VAQNVK", "NOPE-V-VZFLF", "VOZ-FQ-FREIVPR QRFX", "VOZ-NC-FREIVPRABJNQZVAF"}
	var prbStatus = []string{"APPROVED", "CANCELLED", "CLOSED", "ON HOLD", "QUEUED", "RCA COMPLE", "READY FOR", "WORK IN PR"}

	flag := true
	for gbSize < 10 {
		csvData3prb = nil
		csvData3prb = append(csvData3prb, prbStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//ABSTRACT
			tmpSlice = append(tmpSlice, "")

			//rand  ACT_ST_DTTM part1
			timeStampString, timeStamp, _ := randate()

			//rand ACT_FINISH_DTTM
			min := 24
			max := 100
			randFin := rand.Intn(max-min+1) + min
			timeStampFin := timeStamp.Add(time.Duration(randFin) * time.Hour)
			timeStampA := timeStampFin.Format("Jan 02, 2006")
			timeStampB := timeStampFin.Format("15:04:05")
			timeStampFinABString := timeStampA + " @ " + timeStampB + ".000"
			tmpSlice = append(tmpSlice, timeStampFinABString)

			//rand  ACT_ST_DTTM part2
			tmpSlice = append(tmpSlice, timeStampString)

			//AGEBIN
			tmpSlice = append(tmpSlice, "")

			//rand BACID
			min = 0020000
			max = 9999999
			randClientID := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "BAC"+strconv.Itoa(randClientID))

			//rand BACKLOG
			min = 0
			max = 1
			randNum := rand.Intn(max-min+1) + min
			//fmt.Println(randBacklog)
			tmpSlice = append(tmpSlice, strconv.Itoa(randNum))

			//BACKLOG PERCENTAGE
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "100")
			} else {
				tmpSlice = append(tmpSlice, "0")
			}

			//BIN
			tmpSlice = append(tmpSlice, "")

			//rand CALL_CD
			tmpSlice = append(tmpSlice, prbCallCD[rand.Intn(len(prbCallCD))])

			//rand CANCELLED_DF
			tmpSlice = append(tmpSlice, "0")

			//rand CATEGORY
			tmpSlice = append(tmpSlice, prbCategory[rand.Intn(len(prbCategory))])

			//rand CAUSE
			tmpSlice = append(tmpSlice, prbCause[rand.Intn(len(prbCause))])

			//Cl
			tmpSlice = append(tmpSlice, "")

			//Closedticket, CLOSED_DATE_DF,CLOSED_DF,CLOSED_DTTM
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "0")
				tmpSlice = append(tmpSlice, "0")
				tmpSlice = append(tmpSlice, "0")
				tmpSlice = append(tmpSlice, "")
			} else {
				tmpSlice = append(tmpSlice, "1")
				tmpSlice = append(tmpSlice, "1")
				tmpSlice = append(tmpSlice, "1")
				tmpSlice = append(tmpSlice, timeStampFinABString)
			}

			//CLOSE_CODE
			tmpSlice = append(tmpSlice, prbCloseCode[rand.Intn(len(prbCloseCode))])

			//COUNT
			tmpSlice = append(tmpSlice, "1")

			//rand  CREATIONDATE_DF
			timeStampStringDF, timeStampDF, _ := randate()
			tmpSlice = append(tmpSlice, timeStampStringDF)

			//rand DAY_NAME_CREATED
			tmpSlice = append(tmpSlice, strings.ToLower(timeStampDF.Weekday().String()))

			//RESOLVED_DTTM part1
			var timeStampRes time.Time
			timeStampResABString := ""
			if randNum != 1 {
				minRes := 24
				maxRes := 256
				randRes := rand.Intn(maxRes-minRes+1) + minRes
				timeStampRes = timeStamp.Add(time.Duration(randRes) * time.Hour)
				timeStampA := timeStampRes.Format("Jan 02, 2006")
				timeStampB := timeStampRes.Format("15:04:05")
				timeStampResABString = timeStampA + " @ " + timeStampB + ".000"
			}

			//DAY_NAME_RESOLVED
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "")
			} else {
				resolvedWeek := strings.ToLower(timeStampRes.Weekday().String())
				tmpSlice = append(tmpSlice, string(resolvedWeek[0])+string(resolvedWeek[1])+string(resolvedWeek[2]))
			}

			//rand DV_SEVERITY
			min = 1
			max = 5
			oriSeverity := rand.Intn(max-min+1) + min
			tmpSlice = append(tmpSlice, strconv.Itoa(oriSeverity))

			//D_CLIENT_ID
			tmpSlice = append(tmpSlice, strconv.Itoa(randClientID))

			//EFFICIENCYPERCENT
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "0")
			} else {
				tmpSlice = append(tmpSlice, "100")
			}

			//FINAL_STATUS
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "Open")
			} else {
				tmpSlice = append(tmpSlice, "Closed")
			}

			//FOLDER_PATH
			tmpSlice = append(tmpSlice, "")

			//HOUR_CREATED, HOUR_RESOLVED
			if randNum == 1 {
				subtractedDuration := timeStampDF.Sub(time.Now())
				subtractedHours := math.Abs(subtractedDuration.Hours())
				tmpSlice = append(tmpSlice, strconv.Itoa(int(subtractedHours)))
				tmpSlice = append(tmpSlice, "")
			} else {
				subtractedDuration := timeStampRes.Sub(timeStampDF)
				subtractedHours := math.Abs(subtractedDuration.Hours())
				tmpSlice = append(tmpSlice, strconv.Itoa(int(subtractedHours)))
				tmpSlice = append(tmpSlice, strconv.Itoa(int(subtractedHours)))
			}

			//ID
			min = 00000000
			max = 99999999
			randID := rand.Intn(max-min+1) + min
			//fmt.Println(randID)
			tmpSlice = append(tmpSlice, strconv.Itoa(randID))

			//INCIDENT_CODE_ID
			min = 0
			max = 2
			randPrefix := rand.Intn(max-min+1) + min
			min = 00000000
			max = 99999999
			randPrefixID := rand.Intn(max-min+1) + min
			if randPrefix == 0 {
				tmpSlice = append(tmpSlice, "CD"+strconv.Itoa(randPrefixID))
			} else if randPrefix == 1 {
				tmpSlice = append(tmpSlice, "CE"+strconv.Itoa(randPrefixID))
			} else {
				min = 0000000
				max = 9999999
				randPrefixID := rand.Intn(max-min+1) + min
				tmpSlice = append(tmpSlice, "CEO"+strconv.Itoa(randPrefixID))
			}

			//LABEL
			min = 0
			max = 1
			randLabel := rand.Intn(max-min+1) + min
			if randLabel == 1 {
				tmpSlice = append(tmpSlice, "Other")
			} else {
				tmpSlice = append(tmpSlice, "")
			}

			//LAST_REFRESH_DTTM
			minRef := 1
			maxRef := 48
			randRef := rand.Intn(maxRef-minRef+1) + minRef
			timeStampRef := timeStamp.Add(time.Duration(randRef) * time.Hour)
			timeStampA = timeStampRef.Format("Jan 02, 2006")
			timeStampB = timeStampRef.Format("15:04:05")
			timeStampRefABString := timeStampA + " @ " + timeStampB + ".000"
			tmpSlice = append(tmpSlice, timeStampRefABString)

			//MODIFY_DTTM
			randMod := rand.Intn(maxRef-minRef+1) + minRef
			timeStampMod := timeStamp.Add(time.Duration(randMod) * time.Hour)
			timeStampA = timeStampMod.Format("Jan 02, 2006")
			timeStampB = timeStampMod.Format("15:04:05")
			timeStampModABString := timeStampA + " @ " + timeStampB + ".000"
			tmpSlice = append(tmpSlice, timeStampModABString)

			//NO_OWNER_DF
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "100")
			} else {
				tmpSlice = append(tmpSlice, "")
			}

			//NO_OWNER_PR
			tmpSlice = append(tmpSlice, "0")

			//OCCURRED_DTTM
			tmpSlice = append(tmpSlice, "")

			//OCCURRED_DTTM
			tmpSlice = append(tmpSlice, "1")

			//OPEN_DTTM
			tmpSlice = append(tmpSlice, timeStampStringDF)

			//ORIG_RECORD
			randOriRecClass := prbOriRecClass[rand.Intn(len(prbOriRecClass))]
			min = 00000000
			max = 99999999
			randRecID := rand.Intn(max-min+1) + min
			tmpSlice = append(tmpSlice, string(randOriRecClass[0])+string(randOriRecClass[1])+strconv.Itoa(randRecID))

			//ORIG_RECORD_CLASS
			tmpSlice = append(tmpSlice, randOriRecClass)

			//OWNER
			tmpSlice = append(tmpSlice, "1")

			//QUEUE_ID
			tmpSlice = append(tmpSlice, prbQueueID[rand.Intn(len(prbQueueID))])

			//RCA_COMPLETE
			tmpSlice = append(tmpSlice, "0")

			//REASSIGNMENTS
			min = 0
			max = 2
			randReAssign := rand.Intn(max-min+1) + min
			tmpSlice = append(tmpSlice, strconv.Itoa(randReAssign))

			//REOPEN_CUSTOM
			tmpSlice = append(tmpSlice, "")

			//RESOLUTION_CODE
			tmpSlice = append(tmpSlice, "")

			//RESOLVED_DTTM
			tmpSlice = append(tmpSlice, timeStampResABString)

			//RESOLVER_NAME
			tmpSlice = append(tmpSlice, "")

			//rand SEVERITY
			min = 1
			max = 4
			randSeverity := rand.Intn(max-min+1) + min
			tmpSlice = append(tmpSlice, strconv.Itoa(randSeverity))

			//SL
			tmpSlice = append(tmpSlice, "")

			//SLA_STATUS
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "Not Met")
			} else {
				tmpSlice = append(tmpSlice, "")
			}

			//SOURCE_TYPE
			tmpSlice = append(tmpSlice, "DataCenter")

			//STATUS
			tmpSlice = append(tmpSlice, prbStatus[rand.Intn(len(prbStatus))])

			//STATUS_DF
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "BACKLOG")
			} else {
				tmpSlice = append(tmpSlice, "CLOSED")
			}

			//SUBCOMPONENT
			tmpSlice = append(tmpSlice, "")

			//rand TENANT_ID part 1
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = prbTenantID[rand.Intn(len(prbTenantID))]
			}

			//S_DATA_SOURCE
			tmpSlice = append(tmpSlice, "ACME"+string(TenantIDRune))

			//rand TENANT_ID part 2
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//TTR_BIN
			tmpSlice = append(tmpSlice, "")

			//TTR_DAYS
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "")
			} else {
				subtractedDuration := timeStampRes.Sub(timeStampDF)
				subtractedDays := math.Abs(subtractedDuration.Hours() / 24)
				tmpSlice = append(tmpSlice, strconv.Itoa(int(math.Round(subtractedDays*100)/100)))
			}

			//YEAR_MONTH_CREATED
			tmpSlice = append(tmpSlice, strings.Replace(timeStampDF.Format("2006-01"), "-", "_", -1))

			//YEAR_MONTH_RESOLVED
			if randNum == 1 {
				tmpSlice = append(tmpSlice, "")
			} else {
				tmpSlice = append(tmpSlice, strings.Replace(timeStampRes.Format("2006-01"), "-", "_", -1))
			}

			//YEAR_WEEK_CREATED
			createdYear := timeStampDF.Year()
			_, createdWeek := timeStampDF.ISOWeek()
			tmpSlice = append(tmpSlice, strconv.Itoa(createdYear)+"_"+strconv.Itoa(createdWeek))

			//create the final slice
			csvData3prb = append(csvData3prb, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData3prb)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData3prb)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/problem_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData3prb)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "problem_temp.csv") + " > " + filepath.Join(currentPath, "problem_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted for problem")
		}
		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData3prb)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "problem_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 20000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 20000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "unified_problem_processed_v2", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-prb.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 20000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput-prb.json", head, eUser, ePassword, elasticClusterIP, "unified_problem_processed_v2")

			//cleanup the file
			cleanup(currentPath, "finalOutput-prb.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput-prb.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput-prb.json", head, eUser, ePassword, elasticClusterIP, "unified_problem_processed_v2")

		//cleanup the file
		cleanup(currentPath, "finalOutput-prb.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/unified_problem_processed_v2/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/unified_problem_processed_v2/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("unified_problem_processed_v2 index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		time.Sleep(2 * time.Second)

		////
		/////
		//////
	}
	defer wg.Done()
}

func dataSetIncident(wg *sync.WaitGroup, gbSize float64, dataSize int, currentPath string, head string, eUser string, ePassword string, elasticClusterIP string) {

	var csvData2 = [][]string{
		{"actionable", "actual_finish_dttm", "actual_start_dttm", "age", "age_bin", "assignment_group", "assignment_group_parent", "autogenerated", "backlog", "backup", "cancelled", "category", "cdi_last_refresh_dttm", "closed_dttm", "company", "contact_type", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "created_day_name", "created_hour", "created_month", "created_week", "cumulative_hold", "data_source", "description", "dv_priority", "dv_status", "hostname", "incident_code_id", "modified_dttm", "mttr_excl_hold", "mttr_incl_hold", "occurred_dttm", "open_dttm", "original_priority", "person_addressline1", "person_building_id", "person_city", "person_country", "person_department", "person_location", "person_region", "priority", "provider", "provider_account", "reassignment_bucket", "reassignment_count", "reopen_count", "resolution", "resolved_day_name", "resolved_dttm", "resolved_hour", "resolved_month", "resolved_week", "server_building_id", "server_capacity", "server_city", "server_country", "server_drive_type", "server_function", "server_region", "server_site_id", "server_street_address", "short_description", "site_id", "source_type", "sso_ticket", "status", "subcategory", "tenant.account_name", "tenant.bac_id", "tenant.bam_id", "tenant.cdir_id", "tenant.chip_id", "tenant.country", "tenant.region", "tenant_id", "ticket_category", "ticket_class", "ticket_classification", "year_month_created", "year_month_resolved"},
	}

	//sample strings inc
	var incStartString = []string{"actionable", "actual_finish_dttm", "actual_start_dttm", "age", "age_bin", "assignment_group", "assignment_group_parent", "autogenerated", "backlog", "backup", "cancelled", "category", "cdi_last_refresh_dttm", "closed_dttm", "company", "contact_type", "context.application", "context.environment", "context.manage", "context.team", "correlation_id", "created_day_name", "created_hour", "created_month", "created_week", "cumulative_hold", "data_source", "description", "dv_priority", "dv_status", "hostname", "incident_code_id", "modified_dttm", "mttr_excl_hold", "mttr_incl_hold", "occurred_dttm", "open_dttm", "original_priority", "person_addressline1", "person_building_id", "person_city", "person_country", "person_department", "person_location", "person_region", "priority", "provider", "provider_account", "reassignment_bucket", "reassignment_count", "reopen_count", "resolution", "resolved_day_name", "resolved_dttm", "resolved_hour", "resolved_month", "resolved_week", "server_building_id", "server_capacity", "server_city", "server_country", "server_drive_type", "server_function", "server_region", "server_site_id", "server_street_address", "short_description", "site_id", "source_type", "sso_ticket", "status", "subcategory", "tenant.account_name", "tenant.bac_id", "tenant.bam_id", "tenant.cdir_id", "tenant.chip_id", "tenant.country", "tenant.region", "tenant_id", "ticket_category", "ticket_class", "ticket_classification", "year_month_created", "year_month_resolved"}
	var incActionable = []string{"Y", "N"}
	var incBackup = []string{"Y", "N"}
	var incSSOTicket = []string{"Y", "N"}
	var incCategory = []string{"UNKNOWN", "DEVELOPMENT", "PRE_PRODUCTION", "DEVELOPMENT", "PRODUCTION", "DEVELOPMENT", "STAGING", "UNKNOWN", "TEST"}
	var incAssignGroup = []string{"HAXABJA", "NOPE-P-ENGRFTNE", "NOPE-P-JVMEFTNE", "NOPE-P-JVMSBTNE", "NOPE-P-YQOJFHCG", "NOPE-V-FDY", "NOPE-V-FGEOBY", "NOPE-V-FLONFR", "NOPE-V-FUNERCG", "NOPE-V-GJFCEBQHPGFHCC", "NOPE-V-GNZFLF", "NOPE-V-HFIVEHF", "NOPE-V-IZJNER", "NOPE-V-JVERYRFF", "NOPE-V-JZGVIBYV", "NOPE-V-MYVAHKFLF", "NOPE-V-NHGBOBY", "NOPE-V-PFBCFVAQ", "NOPE-V-QO7FLOBY", "NOPE-V-QO7QOOBY", "NOPE-V-RKPUNATR", "NOPE-V-VABENPYOBY", "NOPE-V-VAGFZ", "NOPE-V-VAQFBYNEVF", "NOPE-V-VAQJVAGRY", "NOPE-V-VAQNVK", "NOPE-V-VAQYVAHK", "NOPE-V-VOZPYQ-AJ.SVERJNYY", "NOPE-V-VOZPYQ-FN.JVAQBJF", "NOPE-V-VOZPYQ-FN.YVAHK", "NOPE-V-VOZPYQ-IZJNER", "NOPE-V-VOZPYQ-ONPXHC", "NOPE-V-VZFLF", "NOPE-V-ZDOBY", "NOPE-V-ZIFBCFVA", "NOPE-V-ZQNCFCG"}
	var incContactType = []string{"EMAIL", "Event Management", "EVENTMANAGEMENT", "PHONECALL", "SELFSERVICE"}
	var incAutoGen = []string{"Y", "N"}
	var incContextApp = []string{"\"ACME_DC_Application11\",\"ACME_DC_Application30\"", "\"ACME_DC_Application14\",\"ACME_DC_Application30\"", "\"ACME_DC_Application15\",\"ACME_DC_Application30\"", "\"ACME_DC_Application16\",\"ACME_DC_Application30\"", "\"ACME_DC_Application31\",\"ACME_DC_Application41\",\"ACME_DC_Application42\",\"ACME_DC_Application17\"", "\"ACME_DC_Application31\",\"ACME_DC_Application42\",\"ACME_DC_Application17\"", "\"ACME_DC_Application31\",\"ACME_DC_Application43\",\"ACME_DC_Application19\"", "\"ACME_DC_Application31\",\"ACME_DC_Application43\",\"ACME_DC_Application20\"", "\"ACME_DC_Application32\",\"ACME_DC_Application44\",\"ACME_DC_Application21\"", "\"ACME_DC_Application33\",\"ACME_DC_Application44\",\"ACME_DC_Application21\"", "\"ACME_DC_Application34\",\"ACME_DC_Application21\"", "\"ACME_DC_Application34\",\"ACME_DC_Application44\",\"ACME_DC_Application21\"", "\"ACME_DC_Application35\",\"ACME_DC_Application22\"", "\"ACME_DC_Application36\",\"ACME_DC_Application22\"", "\"ACME_DC_Application36\",\"ACME_DC_Application41\",\"ACME_DC_Application22\"", "\"ACME_DC_Application37\",\"ACME_DC_Application24\"", "\"ACME_DC_Application38\",\"ACME_DC_Application24\"", "\"ACME_DC_Application38\",\"ACME_DC_Application25\"", "\"ACME_DC_Application38\",\"ACME_DC_Application41\",\"ACME_DC_Application24\"", "\"ACME_DC_Application39\",\"ACME_DC_Application25\"", "\"ACME_DC_Application39\",\"ACME_DC_Application26\"", "\"ACME_DC_Application39\",\"ACME_DC_Application41\",\"ACME_DC_Application25\"", "\"ACME_DC_Application4\",\"ACME_DC_Application5\",\"ACME_DC_Application27\",\"ACME_DC_Application28\"", "\"ACME_DC_Application40\",\"ACME_DC_Application10\",\"ACME_DC_Application29\"", "\"ACME_DC_Application40\",\"ACME_DC_Application12\",\"ACME_DC_Application30\"", "\"ACME_DC_Application40\",\"ACME_DC_Application41\",\"ACME_DC_Application26\"", "\"ACME_DC_Application40\",\"ACME_DC_Application9\",\"ACME_DC_Application29\"", "\"ACME_DC_Application42\"", "\"ACME_DC_Application6\",\"ACME_DC_Application28\"", "\"ACME_DC_Application8\",\"ACME_DC_Application28\"", "\"ACME_DC_Application9\",\"ACME_DC_Application28\"", ""}
	var incContextEnv = []string{"UNKNOWN", "DEVELOPMENT", "PRE_PRODUCTION", "DEVELOPMENT", "PRODUCTION", "DEVELOPMENT", "STAGING", "UNKNOWN", "TEST"}
	var incTenantID = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var incHostName = []string{"afkznantre", "awccrrfk56", "awccrrfk57", "awccrrfk58", "awccrrfk59", "awcfvasoxcgfz56", "awicegvopyk56", "awicjpherncc56", "cc-ermcnex", "g6cfprapoffdy56", "inosf58", "ipf7", "ipragre-jqp59", "ipragre-qny65", "j59icefueegz56", "j59icefueegz57", "j59icefueica56", "j59icefuencz56", "j59icefueose56", "j59icefueose57", "j59icefueqaf56", "j59icefueqaf57", "j59icefueqbc56", "j59icefueqbc57", "j59icefueryx51", "j59icefueryx56", "j59icefueryxzba56", "j59icefuexo56", "j59icefuexo57", "j59icefueybt56", "j59icefueybt57", "j59icefuezckl56", "j59icefuezy56", "j59iceoxcgfz56", "j59icerpzjro56", "j59icerpzjro57", "j59icerpzncc50", "j59icerpzncc56", "j59icerpzncc57", "j59icerpzncc58", "j59icerpzncc59", "j59icerpzncc65", "j59icerpznrzc50", "j59icerpznrzc51", "j59icerpznrzc56", "j59icerpznrzc57", "j59icerpznrzc58", "j59icerpznrzc59", "j59icerpznrzn56", "j59icerpznrzn57", "j59icerpzpnu56", "j59icerpzpnu57", "j59icerpzqvep56", "j59icerpzqvep57", "j59icerpzrqo56", "j59icerpzrqo57", "j59icerpzrqo58", "j59icerpzrqo59", "j59icjfrpwzc56", "j59icjfuefcir56", "j59icjfuenq56", "j59icjfuenq57", "j59icjfuewzc56", "j59icjrpzngg56", "j59icjrpzwzc56", "j59igefuerqo56", "j59igerpzjro56", "j59igerpzjro57", "j59igerpzncc52", "j59igerpzncc56", "j59igerpzncc57", "j59igerpzncc58", "j59igerpzncc59", "j59igerpznrz50", "j59igerpznrz51", "j59igerpznrz52", "j59igerpznrz53", "j59igerpznrz54", "j59igerpznrz56", "j59igerpznrz57", "j59igerpznrz58", "j59igerpznrz68", "j59igerpzrqo56", "j59igerpzrqo58", "j59igerpzrqo59", "j59igjrpzwzc56", "jud-jvmpbz-ceq7", "jud-jvmpbz-oqp6", "nmcfnipbecqp56", "nmcfnipbecqp57", "o6cfnicebkgnq56", "o6cfnifrcqz56", "o6cfnifrcva56", "o6cfniibygfdy56", "o6cfniieffdy57", "o6cfninccpbe56", "o6cfnisfpncc56", "o6cfnopragjro56", "o6cnierpbancc56", "o6icynopucncc56", "o6ihynopucncc56", "o78mcejvmncc56", "o78mcejvmncc57", "o78mcejvmzba56", "o78mcejvmzba57", "o78mhejvmncc56", "o7cffnozyttgj57", "o7cfnibz0fbn56", "o7cfnibz0fbn57", "o7cfniepapgk57", "o7cfnipfnszna56", "o7cfnipfuejgf56", "o7cfnippngncc56", "o7cfnipzyttgj53", "o7cfnipzyttgj57", "o7cfnohngfdy56", "o7cfnohngjro56", "o7cfnohngncc56", "o7cfnotbz0qo56", "o7cfnotbz0qo57", "o7cfnotbz0qo58", "o7cfnotgjvncc56", "o7cfnotouq56", "o7cfnotzaqncc56", "o7cfnzqzyttgj57", "o7cfongtqftgj68", "o7cfpbeohfncc58", "o7cfpbeuhzfdy56", "o7cfpbeuhzvak56", "o7qfni9fbntom", "o7qfnibz9ecg56", "o7qfnibz9fbn56", "o7qfnibz9fbnm6", "o7qfnibz9fdy56", "o7qfnohngztg56", "o7qfnotbz8qo56", "o7qfnotbz9qo56", "o7qfnotbz9qo57", "o8ccecebben66", "o8ccecebben67", "o8ccecebben68", "o8ccecebncc68", "o8ccecebncc69", "o8ccecnhben56", "o8ccecrzben56", "o8ccegjben56", "o8ccegjben57", "o8ccegjvncc56", "o8ccegjvncc57", "o8ccenqoben56", "o8ccenqoben57", "o8ccepzfben56", "o8ccepzfncc56", "o8cceqclben56", "o8cceqclncc57", "o8cceqclpqz57", "o8cceqfgben56", "o8cceqfgncc56", "o8cceqfoben56", "o8cceqfoncc56", "o8cceqonben56", "o8cceqonben57", "o8cceqonben58", "o8cceqscben56", "o8cceqyaben56", "o8cceqyancc57", "o8cceqyjben56", "o8cceqypben56", "o8ccerzrncc56", "o8ccerzrncc57", "o8ccexnapfr50", "o8ccjcnhfdy56", "o8ccjfnifsopdq6", "o8ccjfnifsoqve7", "o8ccjfnifsosr56", "o8ccjgcpfpp56", "o8ccjjeyfdy50", "o8ccjjeyfdy57", "o8ccjppnencc56", "o8ccjppnencc57", "o8ccjvasrkz50", "o8ccjvasrkz51", "o8ccjvasrkz56", "o8ccjvasrkz57", "o8ccjvasrkz58", "o8ccrbztfs56", "o8ccrbztfs57", "o8ccrbztfu56", "o8ccrbztfu57", "o8ccrbztfu58", "o8ccrecgrfk56", "o8ccrecgrfk57", "o8ccrfsrfk50", "o8ccrfsrfk51", "o8ccrfsrfk52", "o8ccrfsrfk53", "o8ccrfsrfk56", "o8ccrfsrfk57", "o8ccrfsrfk58", "o8ccrfsrfk59", "o8ccrwhzcfu57", "o8ccrzqzrfk56", "o8ccrzqzrfk57", "o8cfjjeyfdy56", "o8cfjjeyfdy57", "o8cfjjeyfdy59", "o8cfnibz0qgpby6", "o8cfnibz0zd57", "o8cfnibz0zqz57", "o8cfnic3nr56", "o8cfnic3nr57", "o8cfnic3pc56", "o8cfnifzfben56", "o8cfnifzfncc56", "o8cfnigecben56", "o8cfniic3jro56", "o8cfniic3jro57", "o8cfnijngjf57", "o8cfnipxnapfr56", "o8cfnipynben56", "o8cfnjmyqncc57", "o8cfnjmyqncc58", "o8cfnjvmjf57", "o8cfnjvmjf58", "o8cfnjvmjf59", "o8cfnotfgtncc56", "o8cfnotfgtqo56", "o8cfnotvas56", "o8checebben66", "o8checebben67", "o8checebben68", "o8checebncc60", "o8checebncc66", "o8checebncc67", "o8checebncc68", "o8checebncc69", "o8checnhben56", "o8checrzben56", "o8cheenhncc56", "o8chenebflo68", "o8chenebncc66", "o8chenebncc68", "o8chenebncc69", "o8chenqoben56", "o8chenqoben57", "o8chepzfben56", "o8chepzfncc56", "o8cheqanpqz51", "o8cheqanpqz53", "o8cheqanpqz58", "o8cheqclben56", "o8cheqclncc57", "o8cheqclpqz56", "o8cheqfgben56", "o8cheqfoben56", "o8cheqfoncc57", "o8cheqonben56", "o8cheqonben57", "o8cheqscazd56", "o8cheqscben56", "o8cheqyaben56", "o8cheqyjben56", "o8cheqyjncc56", "o8cheqyjncc57", "o8cheqypben56", "o8cheqypncc56", "o8cheqypncc57", "o8cherzrncc56", "o8chjgnfncc56", "o8chjjeyfdy56", "o8chjjeyfdy57", "o8chjqclfdy56", "o8chrbztfs56", "o8chrbztfs57", "o8chrzqzrfk56", "o8cjvasfxyz58", "o8cjvasfxyz59", "o8cjylpcby50", "o8cjylpcby59", "o8cjylppzf50", "o8cjylppzf59", "o8cqenebflo66", "o8cqepynben56", "o8cqeqscazd57", "o8ffprqjc-qo6", "o8ffprqjc-qo7", "o8hfnic3nr56", "o8hfnic3nr57", "o8hfnic3pc56", "o8hfnic3pc57", "o8hfnic3pyqz56", "o8hfnic3pyqz57", "o8hfnifzfben56", "o8hfnifzfncc56", "o8hfniic3fdy56", "o8hfniic3jro56", "o8hfnipzfben56", "o8hfnipzffdy56", "o8icebztfvzncc56", "o8icebztncc50", "o8icebztncc51", "o8icebztncc56", "o8icebztncc57", "o8icebztncc58", "o8icebztncc59", "o8icepngqon56", "o8icepngqon57", "o8icjavegnc56", "o8icjavprnc56", "o8icjavprqo56", "o8icjavprratr56", "o8icjbzincc58", "o8icjcnlyffdy56", "o8icjecgjfncc56", "o8icjecgjfncc57", "o8icjecgjfncc58", "o8icjecgjfncc59", "o8icjecgpzncc56", "o8icjecgpzncc57", "o8icjgrz56", "o8icjjeyjro54", "o8icjnepuefei56", "o8icjnepuejro56", "o8icjnoorffdy57", "o8icjnotbbqtc53", "o8icjnotbbqtp52", "o8icjnotcegt56", "o8icjnotpgkkn50", "o8icjnotpgkkn51", "o8icjpgkkra50", "o8icjpgkkra51", "o8icjpgkkra52", "o8icjpgkkra54", "o8icjpgkkra56", "o8icjpgkqo56", "o8icjpnvff56", "o8icjpoffdy56", "o8icjqnvj57", "o8icjsbqscy56", "o8icjvasacf56", "o8icjvasirnz56", "o8icjvasjro56a", "o8icjvasjro57", "o8icjvasrkqnt56", "o8icjvasrkrqt59", "o8icjwhzcubfg56", "o8icjwhzcubfg57", "o8icyfrpfpna58", "o8ifjjeyjro56", "o8ifjjeyjro57", "o8ihebztncc56", "o8ihebztncc57", "o8ihebztyqnc56", "o8ihezqztoy56", "o8ihjhvcfdy56", "o8ihjhvcjro56", "o8ihjhvcjro57", "o8ihjjeyncc56", "o8ihjnotbbqof50", "o8ihjnotbbqtc57", "o8ihjnotbbqtp56", "o8ihjnotbbqtz58", "o8ihjpbecqp56", "o8ihjylpncc56", "o8ihjylpncc57", "o8iqebztncc56", "o8iqeqonben56", "o8iqjdnncc56", "o8iqjdnqo56", "o8iqjohffdy-g", "o8iqjohffdy56", "o8iqjzyttgj56", "o8qfnibz9hvn56", "o8qfnibz9zqz56", "o8qfniegpben56", "o8qfnigvzncc56", "o8qfnipynben56", "o8qfnipzoncc56", "o8qfnotqrincc56", "o8qfnotvas56", "o8qfnotvas57", "o8qfnotvasyqz56", "o8qjfcecg56", "o8tmcfniperi56", "o8tmcfniperi57", "oc6cfvasrfk56", "oc6cfvasrfk57", "ocvcfvasfwf56", "oemcfninebfym56", "oemcfnipfueqon56", "oemcfnipfuezfp56", "oemcfniplyqqo56", "ofmdfnipfuezfp56", "ofmdfnipfuezfp57", "ofmjepqri57", "ofmqfnipfueqri56", "ofmqfnipfueqri59", "ofmqfniprqjqri56", "ojud-ebire-jjxf", "onopejf050", "onopenc053", "onopenc065", "onopenc066", "onopenc553", "onopenc554", "onopenc561", "onopenc562", "onopenc564", "onopenc565", "onopenc567", "onopenc568", "onopenc569", "onopenc575", "onopenc576", "oqcvasgqztgjl56", "or7qfnisri8frp58", "otmdfnipfuefei56", "otmqfnipfuefei56", "ov6cfvasip56", "ov6cfvasipqo56", "ovcftevaspgk56", "ovcftevaspgk57", "ovcftevaspgk58", "ovcfvasavz96", "ovcfvasegqp57", "ovcfvasfucjro56", "ovcfvasfucncc56", "ovcfvasgfz-yna-serr", "ovcfvasgqz56", "ovcfvasjvaf56", "ovcfvasqqp57", "ovcvasgqztgjl56", "ovqppraqnagqzm7", "ovqppraqnagqzm8", "ovqzmvfn57", "ovvcfqupc59", "pnopenc053", "pnopenc054", "pnopenc060", "pnopenc066", "pnopenc068", "pnopenc069", "pnopenc553", "pnopenc554", "pnopenc561", "pnopenc564", "pnopenc575", "pnopeqo050", "pnopeqo059", "pnopeqo550", "pnopeqo551", "pnopeqo559", "q65icefueegz56", "q65icefueegz57", "q65icefueica56", "q65icefuencz56", "q65icefuencz57", "q65icefueose56", "q65icefueose57", "q65icefueqaf56", "q65icefueqaf57", "q65icefueqbc56", "q65icefueybt56", "q65icefueybt57", "q65icefuezckl56", "q65iceoxcgfz56", "q65icerpzjns57", "q65icerpzjro56", "q65icerpzjro57", "q65icerpzncc50", "q65icerpzncc51", "q65icerpzncc52", "q65icerpzncc53", "q65icerpzncc54", "q65icerpzncc56", "q65icerpzncc57", "q65icerpzncc58", "q65icerpzncc59", "q65icerpzncc65", "q65icerpznczp56", "q65icerpznrzc50", "q65icerpznrzc51", "q65icerpznrzc56", "q65icerpznrzc57", "q65icerpznrzc58", "q65icerpznrzc59", "q65icerpzpnu56", "q65icerpzpnu57", "q65icerpzqvep56", "q65icerpzqvep57", "q65icerpzrqo50", "q65icerpzrqo56", "q65icerpzrqo57", "q65icerpzrqo58", "q65icerpzrqo59", "q65icjfrpwzc56", "q65icjfuefcir56", "q65icjfuenq56", "q65icjfuenq57", "q65icjfueqnvc56", "q65icjfueqnvc57", "q65icjfuewzc56", "q65icjrpzwzc56", "q65igerpzjns56", "q65igerpzjro56", "q65igerpzjro57", "q65igerpzncc56", "q65igerpzncc57", "q65igerpzncc58", "q65igerpzncc59", "q65igerpznczp56", "q65igerpznrz50", "q65igerpznrz51", "q65igerpznrz52", "q65igerpznrz53", "q65igerpznrz54", "q65igerpznrz56", "q65igerpznrz57", "q65igerpznrz58", "q65igerpznrz59", "q65igerpznrz65", "q65igerpznrz66", "q65igerpznrz67", "q65igerpzrqo50", "q65igerpzrqo56", "q65igerpzrqo57", "q65igerpzrqo58", "q65igerpzrqo59", "q65igjrpzwzc56", "qt9bfc7e6iw", "r7cfnipfvtpnc56", "r7cfnipfvtpnc57", "r7cfnipoymqri59", "r7cfnipsygjgf57", "r7cfnipvasuvf54", "r7cfnipvasuvf66", "r7cfnotbz0ncc56", "r7cfnotbz0ncc57", "r7cfnotbz0vag56", "r7cfnotbz0vag57", "r7cfpbecqzmqp57", "r7hfnipfvtpnc56", "r7qfnipfvtpnc56", "r7qfnisri8jf56", "r7qfnotbz8ncc56", "r7qfnotbz9ncc56", "r7qfnotbz9ncc57", "r7qfnotbz9vag56", "r7qfnotbz9vag57", "r7qfnotyzptgj56", "rfkv0", "rfkv1", "rfkv2", "rfkv5", "rfkv6", "rfkv7", "rfkv8", "rfkv9", "ugqp56-221qp154", "ugqp57-221qp154", "zif6", "zif7"}
	var incStatus = []string{"CANCELLED", "CLOSED", "INPROG", "PENDING", "QUEUED", "REJECTED", "RESOLVCONF", "RESOLVED", "SLAHOLD"}
	var incDvStatus = []string{"CANCELLED", "CLOSED", "OPENED"}
	var incServerCap = []string{"Non-Capacity", "Capacity"}
	var incServerSiteID = []string{"b1psavappcor01", "b1psavproxtad01", "b1psavvoltsql01", "b23zprwizapp01", "b23zprwizapp02", "b23zprwizmon01", "b23zprwizmon02", "b2dsabgmndapp01", "b2dsabgom4db01", "b2dsabgom4db02", "b2dsav4soagbz", "b2dsavfev4sec05", "b2dsavom4soa01", "b2dsavom4soaz1", "b2dsavom4sql01", "b2psabgbhd01", "b2psabgmndapp01", "b2psabgom5db02", "b2psabgom5db03", "b2psaelmlggtw01", "b2psavcsafman01", "b2psavcshrwts01", "b2psavcwowapp01", "b2psavcwowapp02", "b2psavom5soa01", "b2psavom5soa02", "b2psbatgdsgtw04", "b2pscorhumsql-t", "b2pscorhumsql01", "b2pswspmlggtw01", "b3dsabgdevapp01", "b3dsavvp8web01", "b3gzpsavcrev01", "b3gzpsavcrev02", "b3pdrclaora01", "b3ppejumpsh01", "b3ppejumpsh02", "b3ppemdmesx01", "b3ppemdmesx02", "b3ppeomgsh01", "b3ppeomgsh02", "b3ppeomgsh03", "b3pperptesx01", "b3ppesfesx03", "b3ppesfesx06", "b3ppesfesx07", "b3pprdbaora01", "b3pprdbaora02", "b3pprdfpora01", "b3pprdlcora01", "b3pprdlnora01", "b3pprdlwora01", "b3pprdnacdm06", "b3pprdsbora01", "b3pprdstora01", "b3ppremeapp01", "b3pprpauora01", "b3pprpemora01", "b3pprproapp13", "b3pprproora11", "b3pprproora12", "b3pprproora13", "b3pprrauapp01", "b3pprtwiapp02", "b3pprtwora01", "b3pprtwora02", "b3ppwccarapp01", "b3ppwhpmpssql03", "b3ppwinfexm06", "b3ppwscmapp01tp", "b3ppwtpcscc01", "b3ppwwrlsql02", "b3ppwwrlsql03", "b3ppwwrlsql04", "b3psabginf01", "b3psabgstgdb01", "b3psavlycedg02", "b3psavom5mdm02", "b3psavom5uia01", "b3psavp8ae01", "b3psavp8ae02", "b3psavp8cp01", "b3psavsmsapp01", "b3psavsmsora01", "b3psawizws03", "b3psawzldapp02", "b3psawzldapp03", "b3pswwrlsql02", "b3pswwrlsql03", "b3puemdmesx01", "b3puradbora02", "b3purcmsora01", "b3purdbaora01", "b3purdbaora02", "b3purdfpora01", "b3purdpycdm01", "b3purdpyora01", "b3puremeapp01", "b3purpauora01", "b3purproapp12", "b3purproora11", "b3purproora12", "b3purproora13", "b3puwpausql01", "b3puwwrlsql01", "b3puwwrlsql02", "b3sscedwp-db1", "b3sscedwp-db2", "b3usavcmsora01", "b3usavcmssql01", "b3usavsmsapp01", "b3vdrdbaora01", "b3vdromgapp01", "b3vdwbussql-t", "b3vdwbussql01", "b3vdwqaapp01", "b3vdwqadb01", "b3vdwuipsql01", "b3vprdynapp01", "b3vprmdmgbl02", "b3vpromgapp01", "b3vpromgapp02", "b3vpromgapp03", "b3vpromgapp04", "b3vpromgapp05", "b3vpromgapp06", "b3vpromgsimapp01", "b3vpwabbessql01", "b3vpwabbesuem01", "b3vpwabposmgn01", "b3vpwarchrweb01", "b3vpwcaiss01", "b3vpwctxdb01", "b3vpwfodfpl01", "b3vpwinfexdag01", "b3vpwinfexedg03", "b3vpwinfexedg04", "b3vpwinfnps01", "b3vpwinfveam01", "b3vpwnicedb01", "b3vpwniceiti01", "b3vpwpaylssql01", "b3vpwrptcmapp01", "b3vpwrptcmapp02", "b3vpwrptwsapp02", "b3vpwtem01", "b3vpwwrdsls01", "b3vpwwrlweb01", "b3vpwwrlweb02", "b3vpwwrlweb03", "b3vpwwrlweb09", "b3vpwwrlweb10", "b3vpwwrlweb12", "b3vurmdmhub01", "b3vuromgapp01", "b3vuromgapp02", "b3vuromgapp03", "b3vuwabgooddb06", "b3vuwexedge01", "b3vuwuipweb02", "b3vuwwrlapp01", "babcrap008", "babcrap009", "babcrap015", "babcrap016", "babcrap019", "babcrap020", "babcrap021", "babcrap508", "babcrap510", "babcrap511", "babcrdb006", "babcrdb504", "babcrdb505", "bcdwizcombdc", "bdpinftdmgtwy01", "bi1psinfvc01", "bidccendantdmz2", "biipsdhcp04", "bipinftdmgtwy01", "bipinfwizcmdc1", "bipsgrinfctx02", "bipsgrinfctx03", "bipsinfdb11", "bipsinfddc01", "bipsinfddc02", "bipsinftsm-lan-free", "bipsinftsm01", "bipsinftsm02", "bipsinftws11", "bp1psinfesx01", "bp1psinfesx02", "bp1psinfesx03", "bpipsinfsjs01", "brzpsavaroslz01", "brzpsavcshrdba01", "brzpsavcylddb01", "bszdsavcshrdev01", "bszpsavcshrweb01", "bszwrcdev02", "bwhq-rover-wwk2", "bwhq-rover-wwks", "cabcrap008", "cabcrap009", "cabcrap013", "cabcrap019", "cabcrap020", "cabcrap508", "cabcrap515", "cabcrdb006", "cabcrdb504", "cabcrdb505", "d10vprecmwaf02", "d10vpwecmjmp01", "d10vpwsecjmp01", "d10vpwshrjmp01", "d10vpwshrspve01", "d10vtrecmaem07", "d10vtrecmapp01", "d10vtrecmapp03", "d10vtrecmapp04", "d10vtrecmedb01", "d10vtrecmedb04", "d10vtrecmwaf01", "d10vtwecmjmp01", "e2dsabglmcgtw01", "e2dsabgom3app01", "e2dsabgom4app01", "e2dsabgom4app02", "e2psabgaltmgn01", "e2psabghpsweb01", "e2psabgisapxy01", "e2psabgom5app01", "e2psabgom5app02", "e2psabgom5int01", "e2psabgom5int02", "e2psabgom5int03", "e2psabgom5int04", "e2psavcblzapp10", "e2psavcblzdev04", "e2psavcfltwts02", "e2psavcsigcap01", "e2psavcsigcap02", "e2psavfev5sec01", "e2ssabgtimapp01", "e2usavcsigcap02", "mvs1", "mvs2", "njpsavcorpdc02", "njpsavp8is02", "njpsinfbkptsm01", "njpsinfwizdc01", "njvpwcureapp01", "nsxmanager", "pp-rezpark", "reserved", "vbpsavcinftsm01", "vbpsinfwizdc01", "w04vprbkptsm01", "w04vprecmaema01", "w04vprecmaemp01", "w04vprecmaemp03", "w04vprecmapp04", "w04vprecmcah02", "w04vprecmdirc02", "w04vprecmedb02", "w04vprecmweb01", "w04vprshrbfr02", "w04vprshrdns01", "w04vprshrdop01", "w04vprshrdop02", "w04vprshrlog01", "w04vprshrmpxy01", "w04vprshrrtm02", "w04vpwecmjmp01", "w04vpwshrspve01", "w04vtrecmaem08", "w04vtrecmapp01", "w04vtrecmapp03", "w04vtrecmapp04", "whq-wizcom-bdc1", "wrc-wizcom-fin1"}
	var incTicketCateg = []string{"Application", "Backup", "CPU", "Database", "Disk", "Job Abends", "Mail", "Node Down", "Other Automata", "Process", "SAP Batch Jobs", "Server Reboot Event", "Service In Alert", "Space", "Unclassified Actionable"}

	flag := true
	for gbSize < 10 {
		csvData2 = nil
		csvData2 = append(csvData2, incStartString)
		for data := 0; data < dataSize; data++ {

			var tmpSlice []string

			//rand  modified_dttm part1
			timeStampString, timeStamp, _ := randate()
			//rand actionable
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incActionable[rand.Intn(len(incActionable))])

			//rand actual_finish
			ranDate, _, _ := randate()
			tmpSlice = append(tmpSlice, ranDate)

			//rand actual_start_dttm
			ranDate, _, _ = randate()
			tmpSlice = append(tmpSlice, ranDate)

			//rand age
			min := -1
			max := 100
			randAge := rand.Intn(max-min+1) + min
			//fmt.Println(randAge)
			tmpSlice = append(tmpSlice, strconv.Itoa(randAge))

			//rand age_bin
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand assignment_group
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incAssignGroup[rand.Intn(len(incAssignGroup))])

			//rand assignment_group_parent
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand autogenerated
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incAutoGen[rand.Intn(len(incAutoGen))])

			//rand backlog
			min = 0
			max = 1
			randBacklog := rand.Intn(max-min+1) + min
			//fmt.Println(randBacklog)
			tmpSlice = append(tmpSlice, strconv.Itoa(randBacklog))

			//rand backup
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incBackup[rand.Intn(len(incBackup))])

			//rand cancelled
			min = 0
			max = 1
			randCancelled := rand.Intn(max-min+1) + min
			//fmt.Println(randCancelled)
			tmpSlice = append(tmpSlice, strconv.Itoa(randCancelled))

			//rand category
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incCategory[rand.Intn(len(incCategory))])

			//rand cdi_last_refresh_dttm
			ranDate, _, _ = randate()
			tmpSlice = append(tmpSlice, ranDate)

			//rand closed_dttm
			ranDate, _, _ = randate()
			tmpSlice = append(tmpSlice, ranDate)

			//rand company
			tmpSlice = append(tmpSlice, "=-COMPANY")

			//rand contact_type
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incContactType[rand.Intn(len(incContactType))])

			//rand context.application
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 0
			max = 4
			randNum := rand.Intn(max-min+1) + min
			collectiveString := ""
			for trav := 0; trav <= randNum; trav++ {
				collectiveString = collectiveString + incContextApp[rand.Intn(len(incContextApp))]
				if trav != randNum {
					collectiveString = collectiveString + ","
				}
			}
			tmpSlice = append(tmpSlice, "["+collectiveString+"]")

			//rand context.environemnt
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\""+incContextEnv[rand.Intn(len(incContextEnv))]+"\"]")

			//rand context.manage
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "[\"IBM\"]")

			//rand context.team
			min = 1
			max = 9
			randTeam := rand.Intn(max-min+1) + min
			contextTeam := "[\"acme sales demo" + strconv.Itoa(randTeam) + "\"]"
			//fmt.Println(randTeam)
			tmpSlice = append(tmpSlice, contextTeam)

			//rand TENANT_ID part1
			TenantIDRune := make([]rune, 24)
			for i := range TenantIDRune {
				TenantIDRune[i] = incTenantID[rand.Intn(len(incTenantID))]
			}

			//rand correlation_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 00000000
			max = 99999999
			seqno := rand.Intn(max-min+1) + min
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, "IBM-"+string(TenantIDRune)+"-1000000000-"+strconv.Itoa(seqno))

			//rand created_day_name
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			timeA := timeStamp.Format("Mon")
			//fmt.Println(timeA)
			tmpSlice = append(tmpSlice, timeA)

			//rand created_hour
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			timeA = timeStamp.Format("15")
			//fmt.Println(timeA)
			tmpSlice = append(tmpSlice, timeA)

			//rand created_month
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			timeA = timeStamp.Format("010206")
			//fmt.Println(string(timeA[0]) + string(timeA[1]))
			tmpSlice = append(tmpSlice, string(timeA[0])+string(timeA[1]))

			//rand created_week
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			_, weekNum := timeStamp.ISOWeek()
			//fmt.Println(weekNum)
			tmpSlice = append(tmpSlice, strconv.Itoa(weekNum))

			//rand cumulative_hold
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "0")

			//rand data_source
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "ACME"+string(TenantIDRune))

			//rand description
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand dv_priority
			min = 1
			max = 5
			randDvPriority := rand.Intn(max-min+1) + min
			//fmt.Println(randDvPriority)
			tmpSlice = append(tmpSlice, strconv.Itoa(randDvPriority))

			//rand dvStatus
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incDvStatus[rand.Intn(len(incDvStatus))])

			//rand hostname
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incHostName[rand.Intn(len(incHostName))])

			//rand incident_code_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 000000000
			max = 999999999
			incCodeID := rand.Intn(max-min+1) + min
			//fmt.Println(incCodeID)
			tmpSlice = append(tmpSlice, "IN"+strconv.Itoa(incCodeID))

			//rand modified_dttm part2
			min = 1
			max = 20
			randMod := rand.Intn(max-min+1) + min
			timeStampMod := timeStamp.Add(time.Duration(randMod) * time.Minute)
			timeStampA := timeStampMod.Format("Jan 02, 2006")
			timeStampB := timeStampMod.Format("15:04:05")
			//timeC := timeStampMod.UnixNano()
			//timeStampInt64str := strconv.FormatInt(timeC, 10)
			//tmpSlice = append(tmpSlice, timeStampInt64str)
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")

			//rand mttr_excl_hold
			min = 1
			max = 99
			randMttrExcl := rand.Intn(max-min+1) + min
			//fmt.Println(randMttrExcl)
			tmpSlice = append(tmpSlice, strconv.Itoa(randMttrExcl))

			//rand mttr_incl_hold
			tmpSlice = append(tmpSlice, strconv.Itoa(randMttrExcl))

			//rand occurred_dttm
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "0")

			//rand open_dttm part2
			tmpSlice = append(tmpSlice, timeStampString)
			//tmpSlice = append(tmpSlice, timeStampString)

			//rand original_priority
			min = 1
			max = 5
			randOriPriority := rand.Intn(max-min+1) + min
			//fmt.Println(randOriPriority)
			tmpSlice = append(tmpSlice, strconv.Itoa(randOriPriority))

			//rand person_addressline1
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_building_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_city
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_country
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_department
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_location
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand person_region
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand priority
			min = 1
			max = 5
			randPriority := rand.Intn(max-min+1) + min
			//fmt.Println(randPriority)
			tmpSlice = append(tmpSlice, strconv.Itoa(randPriority))

			//rand provider
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "IBM DC")

			//rand provider_account
			min = 0200000
			max = 9999999
			randClientID := rand.Intn(max-min+1) + min
			bacid := "BAC" + strconv.Itoa(randClientID)
			//fmt.Println(randClientID)
			tmpSlice = append(tmpSlice, bacid)

			//rand reassignment_bucket
			min = 1
			max = 9
			randReAssign := rand.Intn(max-min+1) + min
			//fmt.Println(randReAssign)
			if randReAssign >= 0 && randReAssign <= 5 {
				tmpSlice = append(tmpSlice, "0 to 5 Times")
			} else if randReAssign >= 6 && randReAssign <= 9 {
				tmpSlice = append(tmpSlice, "6 to 10 Times")
			}

			//rand reassignment_count
			tmpSlice = append(tmpSlice, strconv.Itoa(randReAssign))

			//rand reopen_count
			tmpSlice = append(tmpSlice, "0")

			//rand resolution
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand resolved_day_name
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			min = 24
			max = 128
			randRes := rand.Intn(max-min+1) + min
			timeStampRes := timeStamp.Add(time.Duration(randRes) * time.Hour)
			timeStampA = timeStampRes.Format("Jan 02, 2006")
			timeStampB = timeStampRes.Format("15:04:05")
			timeA = timeStamp.Format("Mon")
			//fmt.Println(timeA)
			tmpSlice = append(tmpSlice, timeA)

			//rand resolved_dttm
			//timeC = timeStampRes.UnixNano()
			//timeStampInt64str = strconv.FormatInt(timeC, 10)
			tmpSlice = append(tmpSlice, timeStampA+" @ "+timeStampB+".000")
			//tmpSlice = append(tmpSlice, timeStampInt64str)

			//rand resolved_hour
			timeA = timeStampRes.Format("15")
			//fmt.Println(timeA)
			tmpSlice = append(tmpSlice, timeA)

			//rand resolved_month
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			timeA = timeStampRes.Format("010206")
			//fmt.Println(string(timeA[0]) + string(timeA[1]))
			tmpSlice = append(tmpSlice, string(timeA[0])+string(timeA[1]))

			//rand resolved_week
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			_, weekNum = timeStampRes.ISOWeek()
			//fmt.Println(weekNum)
			tmpSlice = append(tmpSlice, strconv.Itoa(weekNum))

			//rand server_building_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_capacity
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incServerCap[rand.Intn(len(incServerCap))])

			//rand server_city
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_country
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_drive_type
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_function
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_region
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand server_site_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incServerSiteID[rand.Intn(len(incServerSiteID))])

			//rand server_street_address
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand short_description
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand site_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand source_type
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "DataCenter")

			//rand sso_ticket
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incSSOTicket[rand.Intn(len(incSSOTicket))])

			//rand status
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, incStatus[rand.Intn(len(incStatus))])

			//rand subcategory
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand tenant.account_name
			tmpSlice = append(tmpSlice, contextTeam)

			//rand tenant.bac_id
			tmpSlice = append(tmpSlice, bacid)

			//rand tenant.bam_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "")

			//rand tenant.cdir_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "CDIR-0000000000")

			//rand tenant.chip_id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "100000000")

			//rand tenant.country
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "United States")

			//rand tenant.region
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "AG")

			//rand tenant.id
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, string(TenantIDRune))

			//rand ticket_category
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			ticketCategory := incTicketCateg[rand.Intn(len(incTicketCateg))]
			tmpSlice = append(tmpSlice, ticketCategory)

			//rand ticket_class
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, "INCIDENT")

			//rand ticket_classification
			//fmt.Println(dataSource[rand.Intn(len(dataSource))])
			tmpSlice = append(tmpSlice, ticketCategory+" Issues")

			//rand year_month_created
			tmpSlice = append(tmpSlice, timeStamp.Format("2006-01"))

			//rand year_month_resolved
			tmpSlice = append(tmpSlice, timeStampRes.Format("2006-01"))

			////////////
			///////////////
			////////////////////////
			///////////////////////////////
			////////////////////////////////////////
			////////////
			///////////////
			////////////////////////
			///////////////////////////////
			////////////////////////////////////////
			////////////
			///////////////
			////////////////////////
			///////////////////////////////
			////////////////////////////////////////
			////////////
			///////////////
			////////////////////////
			///////////////////////////////
			////////////////////////////////////////
			////////////
			///////////////
			////////////////////////
			///////////////////////////////
			////////////////////////////////////////

			//create the final slice
			csvData2 = append(csvData2, tmpSlice)

			//fmt.Println("\n\n\n\n", csvData2)

			// Get value from cell by given worksheet name and axis.

		}

		//fmt.Println("\n\n\n\n", csvData2)

		// Open the file
		recordFile, err := os.Create("./datafiles-temporary/incident_temp.csv")
		if err != nil {
			fmt.Println("Error while creating the file::", err)
			return
		}

		// Initialize the writer
		writer := csv.NewWriter(recordFile)

		// Write all the records
		err = writer.WriteAll(csvData2)
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		err = recordFile.Close()
		if err != nil {
			fmt.Println("Error while closing the file ::", err)
			return
		}

		//csvtojson output.csv > output.json
		cmd := "csvtojson " + filepath.Join(currentPath, "incident_temp.csv") + " > " + filepath.Join(currentPath, "incident_temp.json")
		_, err = exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File converted for inc")
		}
		//fmt.Println(out)

		//urlsJson, _ := json.Marshal(csvData2)
		//fmt.Println(string(urlsJson))

		jsonOutputFile, err := os.Open(filepath.Join(currentPath, "incident_temp.json"))
		if err != nil {
			fmt.Println(err)
		}

		// read our opened xmlFile as a byte array.
		byteOutputFile, _ := ioutil.ReadAll(jsonOutputFile)

		defer jsonOutputFile.Close()

		counterFloorVal := 0
		counterRemainingPairs := 0
		trackerOutputfile := 0

		var mapOutputFile []map[string]interface{}
		//var finalOutputFile []map[string]interface{}

		json.Unmarshal([]byte(byteOutputFile), &mapOutputFile)

		fmt.Println("This is output file length", len(mapOutputFile))

		var equalPairs float64
		equalPairs = float64(len(mapOutputFile)) / 1000
		//fmt.Println("Number of pairs are", equalPairs)
		equalPairsFloor := math.Floor(equalPairs)
		//equalPairsCeil := math.Ceil(equalPairs)
		//fmt.Println("Floor val", equalPairsFloor)
		//fmt.Println("Ceil val", equalPairsCeil)
		totalNormalPairs := equalPairsFloor * 1000
		//fmt.Println("Total pairs of hundred", totalNormalPairs)
		remainingPairs := float64(len(mapOutputFile)) - totalNormalPairs
		//fmt.Println("Remaining pairs", remainingPairs)

		var m = map[string]interface{}{"index": map[string]interface{}{"_index": "incidents", "_type": "_doc"}}

		for ; counterFloorVal != int(equalPairsFloor); counterFloorVal++ {
			//open the file

			// If the file doesn't exist, create it, or append to the file
			fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			for z := 0; z < 1000; z++ {
				//fmt.Println("Written file at index:", trackerOutputfile)
				writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
				trackerOutputfile++
			}

			//close the file
			if err := fiiiile.Close(); err != nil {
				log.Fatal(err)
			}

			//post the file
			bulkPOST(currentPath, "finalOutput.json", head, eUser, ePassword, elasticClusterIP, "incidents")

			//cleanup the file
			cleanup(currentPath, "finalOutput.json")

			//fmt.Println("counterFloorVal is:", counterFloorVal)

		}

		//post remaining values
		fiiiile, err := os.OpenFile(filepath.Join(currentPath, "finalOutput.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("Time to write the final data")
		for ; counterRemainingPairs < int(remainingPairs); counterRemainingPairs++ {
			writeFinalOutputArc1(fiiiile, m, mapOutputFile[trackerOutputfile])
			//fmt.Println("Written file at index:", trackerOutputfile)
			trackerOutputfile++
			//fmt.Println("counterRemainingPairs is:", counterRemainingPairs)

		}
		//close the file
		if err := fiiiile.Close(); err != nil {
			log.Fatal(err)
		}

		//post the file
		bulkPOST(currentPath, "finalOutput.json", head, eUser, ePassword, elasticClusterIP, "incidents")

		//cleanup the file
		cleanup(currentPath, "finalOutput.json")

		/////////

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		///////////
		if flag == true {

			url := head + elasticClusterIP + ":9200/incidents/_settings"

			strReq := `{"index.mapping.total_fields.limit": 100000}`

			var strBytes = []byte(strReq)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(strBytes))
			if err != nil {
				log.Fatalf("Error Occured in GET for index stats", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(eUser, ePassword)

			response, err := client.Do(req)
			if err != nil && response == nil {
				fmt.Println("Error sending request to API endpoint.", err)
			}

			flag = false
		}

		//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"

		url := head + elasticClusterIP + ":9200/incidents/_stats/store"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error Occured in GET for index stats", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(eUser, ePassword)

		response, err := client.Do(req)
		if err != nil && response == nil {
			fmt.Println("Error sending request to API endpoint.", err)
		}

		body, _ := ioutil.ReadAll(response.Body)
		json.Marshal(body)
		//fmt.Println("response Body:", string(body))
		//fmt.Println("response StatusCode:", response.StatusCode)
		defer response.Body.Close()

		var resIndPList map[string]interface{}
		json.Unmarshal([]byte(body), &resIndPList)
		resVal := resIndPList["_all"].(map[string]interface{})
		resVal2 := resVal["total"].(map[string]interface{})
		resVal3 := resVal2["store"].(map[string]interface{})
		byteSize := resVal3["size_in_bytes"].(float64)
		//byteSize, _ := strconv.Atoi(resVal3["size_in_bytes"].(float64))
		fmt.Println("Incidents index size is")
		fmt.Println(byteSize)
		gbSize = ((byteSize / 1024) / 1024) / 1024
		fmt.Println(gbSize)

		///////////

		fmt.Println("Conversion happened successfully")

		fmt.Println("Time to wait")

		time.Sleep(5 * time.Second)

		////
		/////
		//////
	}
	defer wg.Done()
}

func bulkPOST(currentPath string, fileName string, head string, eUser string, ePassword string, elasticClusterIP string, titleIndP string) {
	//curl -H 'Content-Type: application/x-ndjson' -XPOST 'elk-elastic-service:9200/incident-fv-2/_bulk?pretty' --data-binary @finalOutput.json
	//url := "http://169.63.76.233:32439/" + titleIndP + "/_bulk?pretty"
	//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/" + titleIndP + "/_bulk?pretty"
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	//url := "http://elastic:" + elasticPass + "@" + elasticClusterIP + ":9200/.kibana/_search?size=1000"
	url := head + elasticClusterIP + ":9200/" + titleIndP + "/_bulk?pretty"
	//fmt.Println("The url is", url)

	jsonStr, err := ioutil.ReadFile(filepath.Join(currentPath, fileName))
	if err != nil {
		panic(err)
	}
	//fmt.Println("File is opened")
	//fmt.Println("File is in bytes format")
	//fmt.Println(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-ndjson")
	req.SetBasicAuth(eUser, ePassword)

	res, err := client.Do(req)
	if err != nil && res == nil {
		fmt.Println("Error sending request to API endpoint.", err)
	}

	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println("response Body:", string(body))

	defer res.Body.Close()

	/////////////
	/*url := "http://" + elasticClusterIP + ":9200/incident/_bulk?pretty"
	fmt.Println("The url is", url)

	jsonStr, err := ioutil.ReadFile(filepath.Join(currentPath, "finalOutput.json"))
	if err != nil {
		panic(err)
	}
	fmt.Println("File is opened")
	fmt.Println("File is in bytes format")
	//fmt.Println(jsonStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	fmt.Println("POST loaded")

	req.Header.Set("Content-Type", "application/x-ndjson")
	req.Header.Set("ChannelName", "mcmpadmin")
	req.Header.Set("ChannelPassword", "mcmp@Passwd")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Do request generated")

	defer res.Body.Close()*/

	//fmt.Println("response Status:", res.Status)
	//fmt.Println("response Headers:", res.Header)
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println("response Body:", string(body))

}

func writeFinalOutputArc1(fiiiile *os.File, m map[string]interface{}, val map[string]interface{}) {

	byteFinalOutputFile1, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("Issue with the Dashboard Marshal")
		panic(err)
	}
	byteFinalOutputFile2, err := json.Marshal(&val)
	if err != nil {
		fmt.Printf("Issue with the Dashboard Marshal")
		panic(err)
	}

	if _, err := fiiiile.Write([]byte(string(byteFinalOutputFile1) + "\n")); err != nil {
		log.Fatal(err)
	}
	if _, err := fiiiile.Write([]byte(string(byteFinalOutputFile2) + "\n")); err != nil {
		log.Fatal(err)
	}

}

func randate() (string, time.Time, string) {
	min := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 12, 30, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	timeimte := time.Unix(sec, 0)
	timeA := timeimte.Format("Jan 02, 2006")
	timeB := timeimte.Format("15:04:05")
	timeC := timeimte.UnixNano()
	timeStampInt64str := strconv.FormatInt(timeC, 10)
	return timeA + " @ " + timeB + ".000", timeimte, timeStampInt64str
}

func cleanup(currentPath string, fileName string) {
	//cleanup
	cmd := "rm " + filepath.Join(currentPath, fileName)
	//fmt.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(out)

}
