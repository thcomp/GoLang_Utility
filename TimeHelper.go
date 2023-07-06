// this file is generated by command.
package utility

import (
    "time"
)

const TzTextAsiaTokyo = "Asia/Tokyo"

type TimeHelper struct {
    tzLocationAsiaTokyo *time.Location
}

var sInsTimeHelper *TimeHelper

func (ins *TimeHelper) DateInAsiaTokyo(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if ins.tzLocationAsiaTokyo == nil {
        if tempTz, loadErr := time.LoadLocation(TzTextAsiaTokyo); loadErr == nil {
            ins.tzLocationAsiaTokyo = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocationAsiaTokyo != nil {
        ret = time.Date(year, month, day, hour, min, sec, nsec, ins.tzLocationAsiaTokyo)
    }

    return
}

func DateInAsiaTokyo(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.DateInAsiaTokyo(year, month, day, hour, min, sec, nsec)
}

func (ins *TimeHelper) NowInAsiaTokyo() (ret time.Time, retErr error) {
    if ins.tzLocationAsiaTokyo == nil {
        if tempTz, loadErr := time.LoadLocation(TzTextAsiaTokyo); loadErr == nil {
            ins.tzLocationAsiaTokyo = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocationAsiaTokyo != nil {
        ret = time.Now().In(ins.tzLocationAsiaTokyo)
    }

    return
}

func NowInAsiaTokyo() (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.NowInAsiaTokyo()
}

func (ins *TimeHelper) InAsiaTokyo(srcTime time.Time) (ret time.Time, retErr error) {
    if ins.tzLocationAsiaTokyo == nil {
        if tempTz, loadErr := time.LoadLocation(TzTextAsiaTokyo); loadErr == nil {
            ins.tzLocationAsiaTokyo = tempTz
        } else {
            retErr = loadErr
        }
    }

    if ins.tzLocationAsiaTokyo != nil {
        ret = srcTime.In(ins.tzLocationAsiaTokyo)
    }

    return
}

func InAsiaTokyo(srcTime time.Time) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.InAsiaTokyo(srcTime)
}

func (ins *TimeHelper) DateInJST(year int, month time.Month, day, hour, min, sec, nsec int) (time.Time, error) {
    return ins.DateInAsiaTokyo(year, month, day, hour, min, sec, nsec)
}

func DateInJST(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.DateInJST(year, month, day, hour, min, sec, nsec)
}

func (ins *TimeHelper) NowInJST() (time.Time, error) {
    return ins.NowInAsiaTokyo()
}

func NowInJST() (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.NowInJST()
}

func (ins *TimeHelper) InJST(srcTime time.Time) (time.Time, error) {
    return ins.InAsiaTokyo(srcTime)
}

func InJST(srcTime time.Time) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.InJST(srcTime)
}


func (ins *TimeHelper) DateInUTC(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    ret = time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
    

    return
}

func DateInUTC(year int, month time.Month, day, hour, min, sec, nsec int) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.DateInUTC(year, month, day, hour, min, sec, nsec)
}

func (ins *TimeHelper) NowInUTC() (ret time.Time, retErr error) {
    ret = time.Now().UTC()
    

    return
}

func NowInUTC() (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.NowInUTC()
}

func (ins *TimeHelper) InUTC(srcTime time.Time) (ret time.Time, retErr error) {
    ret = srcTime.UTC()
    

    return
}

func InUTC(srcTime time.Time) (ret time.Time, retErr error) {
    if sInsTimeHelper == nil {
        sInsTimeHelper = &TimeHelper{}
    }

    return sInsTimeHelper.InUTC(srcTime)
}

