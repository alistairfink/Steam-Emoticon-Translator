package JsonObjects

import (

)

type Alphabet struct {
	Letters []Letter `json:"letters"`
}

type Letter struct {
	Letter string `json:"letter"`
	Emoticons []string `json:"emoticons"`
}