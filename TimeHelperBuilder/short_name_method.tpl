func (ins *{{.Structure}}) NowIn{{.ShortName}}() (time.Time, error) {
    return ins.NowIn{{.Region}}{{.City}}()
}

func NowIn{{.ShortName}}() (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.NowIn{{.ShortName}}()
}
