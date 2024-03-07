func (ins *{{.Structure}}) Now(tzText string) (ret time.Time, retErr error) {
    ins.createTzTextNowMap()

    lowerTzText := strings.ToLower(tzText)
    if targetNowFunc,exist := ins.tzTextNowMap[lowerTzText]; exist {
        ret, retErr = targetNowFunc()
    } else {
        retErr = fmt.Errorf("unknown timezone text: %s", tzText)
    }

    return
}

func Now(tzText string) (ret time.Time, retErr error) {
    if sIns{{.Structure}} == nil {
        sIns{{.Structure}} = &{{.Structure}}{}
    }

    return sIns{{.Structure}}.Now(tzText)
}

func (ins *{{.Structure}}) createTzTextNowMap() {
    if ins.tzTextNowMap == nil {
        ins.tzTextNowMap = map[string](func()(time.Time, error)){
            {{.TzTextNowMap}}
        }
    }
}
