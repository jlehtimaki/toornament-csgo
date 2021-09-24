package toornament

import "fmt"

func GetSeed() ([]byte, error) {
	getGroups()
	return nil, nil
}

func getGroups()([]byte, error){
	subPath := fmt.Sprintf("organizer/v2/groups")
	rangeKey := "groups=0-49"

	data, err := toornamentRest(subPath, rangeKey)
	if err != nil { return nil, err }

	fmt.Println(data)
	return nil, nil
}

func getGroup(){

}