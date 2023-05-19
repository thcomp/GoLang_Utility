func (ins *{{.Structure}}) NowIn{{.Region}}{{.City}}() (ret time.Time, retErr error) {
    {{if eq .City ""}} ret = time.Now().UTC()
    {{ else }}if ins.tzLocation{{.Region}}{{.City}} == nil {
        if tempTz, loadErr := time.LoadLocation(TzText{{.Region}}{{.City}}); loadErr == nil {
            ins.tzLocation{{.Region}}{{.City}} = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocation{{.Region}}{{.City}} != nil {
        ret = time.Now().In(ins.tzLocation{{.Region}}{{.City}})
    }{{end}}

    return
}

func NowIn{{.Region}}{{.City}}() (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.NowIn{{.Region}}{{.City}}()
}
