func (ins *{{.Structure}}) NowIn{{.ShortName}}() (time.Time, error) {
    return ins.NowIn{{.Region}}{{.City}}()
}
