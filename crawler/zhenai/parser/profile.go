package parser

import (
	"goLearn/crawler/engine"
	"goLearn/crawler/model"
	"regexp"
	"strconv"
)

//预先编译好要提取信息的正则表达式
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var residenceRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Residence = extractString(contents, residenceRe)
	profile.House = extractString(contents, houseRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

//封装正则匹配函数,提取字符串
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
