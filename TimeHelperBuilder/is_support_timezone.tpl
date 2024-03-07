func IsSupportTimezone(timezoneText string) (ret bool) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.IsSupportTimezone(timezoneText)
}

func (ins *{{.Structure}}) IsSupportTimezone(timezoneText string) (ret bool) {
    supportedRegionCityNames := []string{
        {{range $index, $regioncityname := .RegionCityName}}"{{$regioncityname}}",{{end}}
    }

    lowerTimezoneText := strings.ToLower(timezoneText)
    for _, supportedRegionCityName := range supportedRegionCityNames {
        if supportedRegionCityName == lowerTimezoneText {
            ret = true
            break
        }
    }

    return
}
