func (ins *{{.Structure}}) DateIn{{.ShortName}}(year int, month time.Month, day, hour, min, sec, nsec int) (time.Time, error) {
    return ins.DateIn{{.Region}}{{.City}}(year, month, day, hour, min, sec, nsec)
}

func DateIn{{.ShortName}}(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.DateIn{{.ShortName}}(year, month, day, hour, min, sec, nsec)
}

func (ins *{{.Structure}}) NowIn{{.ShortName}}() (time.Time, error) {
    return ins.NowIn{{.Region}}{{.City}}()
}

func NowIn{{.ShortName}}() (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.NowIn{{.ShortName}}()
}

func (ins *{{.Structure}}) In{{.ShortName}}(srcTime time.Time) (time.Time, error) {
    return ins.In{{.Region}}{{.City}}(srcTime)
}

func In{{.ShortName}}(srcTime time.Time) (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.In{{.ShortName}}(srcTime)
}

