func (ins *{{.Structure}}) DateIn{{.Region}}{{.City}}(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    {{if eq .City ""}}ret = time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
    {{ else }}if ins.tzLocation{{.Region}}{{.City}} == nil {
        if tempTz, loadErr := time.LoadLocation(TzText{{.Region}}{{.City}}); loadErr == nil {
            ins.tzLocation{{.Region}}{{.City}} = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocation{{.Region}}{{.City}} != nil {
        ret = time.Date(year, month, day, hour, min, sec, nsec, ins.tzLocation{{.Region}}{{.City}})
    }{{end}}

    return
}

func DateIn{{.Region}}{{.City}}(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.DateIn{{.Region}}{{.City}}(year, month, day, hour, min, sec, nsec)
}

func (ins *{{.Structure}}) NowIn{{.Region}}{{.City}}() (ret time.Time, retErr error) {
    {{if eq .City ""}}ret = time.Now().UTC()
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

func (ins *{{.Structure}}) In{{.Region}}{{.City}}(srcTime time.Time) (ret time.Time, retErr error) {
    {{if eq .City ""}}ret = srcTime.UTC()
    {{ else }}if ins.tzLocation{{.Region}}{{.City}} == nil {
        if tempTz, loadErr := time.LoadLocation(TzText{{.Region}}{{.City}}); loadErr == nil {
            ins.tzLocation{{.Region}}{{.City}} = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocation{{.Region}}{{.City}} != nil {
        ret = srcTime.In(ins.tzLocation{{.Region}}{{.City}})
    }{{end}}

    return
}

func In{{.Region}}{{.City}}(srcTime time.Time) (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.In{{.Region}}{{.City}}(srcTime)
}
