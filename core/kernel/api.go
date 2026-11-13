package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func YBI53amDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0iMhlbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1Itc1wWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRquoTymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQhw2m9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T84S4AFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hxk0s0zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGIKXqEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXIVJ8UYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeFkVMIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubHcaUEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2yZ1Uot9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKVsfYATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWpcLMoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6OLaLVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jp1WOB71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSdp4wyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbaV7HJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSCz7vCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gD3PD3S5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwT5aVw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LgBwuI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrjSuR9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdLc5y3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtPXYfaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mebv5gN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDWH4TdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVmG50TgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7NL5xs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoM0UR0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SSY1PUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fo2FYaKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZY6YsDpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgKAXqq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MepneKe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbRYqMq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KACTBmofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2mwlSfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNiWq9Y3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVDwsgOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSCly5uyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48rRBROVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqsoGqttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fagBdZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQZ0OQWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyrkhULYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfGrTSy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfBfR6jUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CP2COAuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUaB4jvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcq8LV7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGPoHLw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkKIE4tPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3B76ESKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vdz6Qj7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdXhi2CsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMDh5sdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkC0a4xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ps7MOeAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9BFbFrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ziOJ7gIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJshRjeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ks56AXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a69aUPwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MKxHv9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weOZQdE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiOm0KZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8DQewH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpFSfTAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5LB1eowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXJSl4IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITTD1FWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS9WNizzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDnHw2ijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51Td4HFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsi6Mnq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1q5GwJHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoL9jNCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3qor335Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVXsfDlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycT7q3DnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BTcVFWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luDELl1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYWAJ4F6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6gdFeeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOlnuexVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cq8YWUD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSQC35HCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmJZK3cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MQ4bWWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEUCuQhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GN0E2jhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgedS12HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfjTXtwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9VqLmWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDSE5SWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 319WHlSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ctb5Rd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulPXm9WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejxhyaOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsfLNJJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuivsnJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYgqUcxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFyrQbBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0uulIBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkjAPv1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkpZsmZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZ0wFHaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZUQHu1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlOGVPFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEwtK06SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hbs6whnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ShmvxJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGC9T3YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJtOg8LmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pxTv0msWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exEhMxjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goZmbryQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stL4Aq1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpH761FvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWYRz5g3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFh7NyiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OrT9wbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLKqExBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hE3AJXEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQo4fZr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwQ2urDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxUgg56NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cODgYOXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rk9XM36sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pg9GxwL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJza2je1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWth7jwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxwWyI72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHyXNCLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXIHwk5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDbcgXlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZSDykqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oz8M2fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUcUWKTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kA3EivoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HwulfNQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkEzUSxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNfjqpIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGpeiv8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIYwsfInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojGggGkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdT1BV2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KOe8lexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLXRDZmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HJr8YXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l751uk0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QATcI7A6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3z8lCxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onHEthHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i30SzCIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6NuH1lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LrxkEgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrGVpF13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brO947GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6itzuy6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RdXjqDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBbIh3aOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvVelYQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNYgkwnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcoeQF3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mr18Z0rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxK9k7WRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MM8jiEizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPr0F6V5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXSHobsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NphQPbfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47BNGMVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRCZtHofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tJxNbAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDmtOzjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcGtF7HpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xa6YaopMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYUXDMMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D53fnDHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WsO6Qb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Be9I0337Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bpmGxhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5SdjsARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QIqubifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfqlAYAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxVHE4u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePWjaRkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ7ArkSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNrzn9gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMWFqEgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOlnM1HuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pc2TcjbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1vYWWTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2DEleGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFUS3j20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3scFhq4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyOwjMlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhOfUBJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwHpker3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqA3k53KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akeVvMmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J55Kw2TDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWn5r1N3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRtQAtLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISQpaZQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLL1xGkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hD3jwyH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFcce4OfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drzyyYdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWAGAm8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wjux8khiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypZBeg4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uo2NiaTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDniV1G7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDrvy26LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjjlHvN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xB2yusOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dn91Z9QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhqGJXUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dppMdH39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeJZoq6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvLhtQiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnaBWL01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HP32h02OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7heq5xSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjDPDwuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bofFSJGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGmyJWJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFfmXecnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BsmzycUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuBP7LWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGiwWvZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqz4iHJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtG79cDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Om0YIzJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UkJnq62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGGDnfEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELlMEew7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYHhoSLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xxq5I6V1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpHp5rqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdccr8AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2r3eCCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfWffE30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9sEPSJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gum6MkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dYPY3kXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbBoqOB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grKLPeh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xcm4LTfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cD79Y4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWou5WEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ds0gyDqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TevadlBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeXC1DN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZbnvxPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lY2A148zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8Jkvp4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKucX862Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTlufDCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLBRn9YeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyTV5YuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQE3vajcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1vCiPK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzfoRII8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddqArDXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Buc8m0HTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU53Y6ZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czcES2JPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFgeRNJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLdC8wNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNV6Ml1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJFig9x4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFj3YonFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MEm8YhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhMogXGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzg1yqunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdwo1WpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gne0u9YlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sq73UocGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ar4qhYYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRUZafW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNuYEqLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func We7wcMkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuGpAso2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slHJ9HIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtlUo2jnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvMBoDcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N68Zf5NMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvWlFO6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUbhlKq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flESqAW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eF99gCZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6IUuJmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HIpJGCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTSugyMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiQjeUYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5dNwI7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ONSDdoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjLCSt0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsdZ3YNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvHkLL0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVNtl1PJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbjnxUn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JH3ZV3N2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvvb13enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0i5BoQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GynJ8zGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Blfq3AjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czmNdUmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjCuHwwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ly4HWLqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EopJPbGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqbOROSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jRM9XtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnRZDP4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhGsEJrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8q3qdg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKrcaq8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gh40FbJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbKB4739Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvlYBPQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjiyECiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dhMm5WzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNHXf6RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dirCjXizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyVMOdu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rZkpwgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRC8UKmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJOatN9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gXPL6d7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvLTi6MMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bH5xaxUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoqM4hOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v30wgq27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7CDXKxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9Pcmm4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaI5HvdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rW4I9ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIRRh5ckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBnIa9RVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tukTf3iGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEki8VTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vi4nFoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNttiJuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LIkYjp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hy2ODk1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPZ9FcwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WnBat9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hajPk774Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCOlVzNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSwDDGVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bU6h0LNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPOoGkgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KA3SPN0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7QTsBiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ass0V64mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2XtjjwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vyvoDekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2LbrlUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7vVVCZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Gz7CobfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jgd2AXCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6N4ii1z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nOcumLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yMigXbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRlnWEeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5WhsyFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tblBmxtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEK9KGY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RydQ3g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4LXZv7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyCJfIenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJe3VkvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNtxw9M2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQ20FmlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPDdpf4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REfi2eErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMHM9xRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSMQ3E34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBAzHKTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdOKS7RfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ari66NSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bNEBVqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CwWxELLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vF9cwvdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5z788vp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZB7JIYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7JnXMfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTtgTTgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yu2AVV5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHy9mtmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDiHWQkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSsZMoHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18xeD44ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGLGR3tMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuwAd5jdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAQffqTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gN4lDjH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBnADaFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBpRrILxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZSqVZadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLr0NtLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9oyEcF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhDBXvnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOvtqFJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZTsYxuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7k45LW8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFSLfCcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79mqTx0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMxnck8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IR6mxZFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvJ21tUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0oetNqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6aSUAJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNO7PX6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxxtqlxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0K2v0CHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVEL4HWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8t9mYOXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1E7742qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofqYImUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APUa1wz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmnSZwBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52FhPEJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z84XkTzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HONd7CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W769su94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWJ6ceB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5K4BLNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whqiDr5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJIEKbggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tz99wDBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTLidOt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4APhvlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llnz8unEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zhbrt04bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NvP9hqEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0AQkd6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 368zkQFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2akd6fCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OX1r5tfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXLg0qaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jW4IKu7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6gVIErrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rxs5wcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEyJLrnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DE2cS4HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZOR3uZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceTzZXRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9BjPlskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pIFnFq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y41Fp3WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhLYGYRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQgSHnT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igr6jsCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkggJBFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdwXH4WnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tKAuLAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCv0faxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e97HKcnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djpRavyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6q3vK7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrVVPReJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grBm0mO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVWDoJEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECh6CX1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MA0q6FlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxlPosBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmWDRJUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYnZAS3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yb1PjicLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ8HnNNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJnxwwcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FwS7DuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOiQ4YDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytR97jPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIEUsvTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajijCaWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJgAvanGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I795utbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N03tpyrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcsUsNweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QMmkBNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4re7D3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miSCFeMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtsDZN5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSwHcrksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugcV86nKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hAv2z8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JL9K1xhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8H79Y849Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy3cTdszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vf1NPenqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MasfMNwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajrOd0XzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxkrhKuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4DSaGt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTTAPSvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtlgLXtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbzaL0lWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpWcpFmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLURHmBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgoQOTCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40z0bJn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEE3z5MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LAEYoz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BeMZ0vP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PfH0Q7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeWFPnwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4ZcHEGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3AyLI2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uB946WIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeBOMD1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DqCEYn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX10pDsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5Mo320DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfIDP0nQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZ4nspGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiCKXbZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7rYrHcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DaOjlEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNKoVshzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBAd5k3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

