package utils

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)
// Fungsi ParseBody digunakan untuk membaca dan mengurai body dari permintaan HTTP yang diberikan,
// kemudian memasukkan data hasil parsing ke dalam variabel yang disediakan.
func ParseBody(r *http.Request, x interface{}) {
    // Membaca seluruh isi body dari permintaan HTTP.
    if body, err := ioutil.ReadAll(r.Body); err == nil {
        // Mengurai isi body yang telah dibaca menggunakan format JSON,
        // kemudian memasukkan hasil parsing ke dalam variabel yang disediakan (x).
        if err := json.Unmarshal([]byte(body), x); err != nil {
            return // Menghentikan eksekusi jika terjadi kesalahan saat parsing.
        }
    }
}
