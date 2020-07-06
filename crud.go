package main
 
import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"
 
    _ "github.com/go-sql-driver/mysql"
)
 
//membuat type transaksi dengan struktur
type transaksi struct {
    no_nota string
    jenis   string
    nama    string
    bayar   string
}
 
//membuat type response dengan struktur
type response struct {
    Status bool
    Pesan  string
    Data   []transaksi
}
 
//membuat fungsi koneksi dengan sql
//sintax -> sql.Open("mysql", "user:password@tcp(host:port)/nama_database")
//karena bawaan xampp password kosong jd dikosongkan saja
func koneksi() (*sql.DB, error) {
    db, salahe := sql.Open("mysql", "devi:devi@tcp(127.0.0.1:3306)/ajc")
    if salahe != nil {
        return nil, salahe
    }
    return db, nil
}
 
//fungsi tampil data
func tampil(pesane string) response {
    db, salahe := koneksi()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:"Gagal Koneksi: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
defer db.Close()
dataSup, salahe := db.Query("select * from transaksi")
if salahe != nil {
    return response{
        Status: false,
        Pesan:  "Gagal Query: " + salahe.Error(),
        Data:   []transaksi{},
    }
}
 
defer dataSup.Close()
var hasil []transaksi
for dataSup.Next() {
    var sup = transaksi{}
    var salahe = dataSup.Scan(&sup.no_nota, &sup.jenis, &sup.nama, &sup.bayar)
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Baca: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
        hasil = append(hasil, sup)
    }
 
    salahe = dataSup.Err()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Kesalahan: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    return response{
        Status: true,
        Pesan:  pesane,
        Data:   hasil,
    }
}
 
//fungsi tampil data berdasarkan id
func getSup(no_nota string) response {
    db, salahe := koneksi()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Koneksi: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    defer db.Close()
    dataSup, salahe := db.Query("select * from transaksi where no_nota=?", no_nota)
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Query: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    defer dataSup.Close()
    var hasil []transaksi
    for dataSup.Next() {
        var sup = transaksi{}
        var salahe = dataSup.Scan(&sup.no_nota, &sup.jenis, &sup.nama, &sup.bayar)
        if salahe != nil {
            return response{
                Status: false,
                Pesan:  "Gagal Baca: " + salahe.Error(),
                Data:   []transaksi{},
            }
        }
 
        hasil = append(hasil, sup)
    }
   
    salahe = dataSup.Err()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Kesalahan: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
   
    return response{
        Status: true,
        Pesan:  "Berhasil Tampil",
        Data:   hasil,
    }
}
 
//fungsi tambah data
func tambah(no_nota string, jenis string, nama string, bayar string) response {
    db, salahe := koneksi()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Koneksi: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    defer db.Close()
    _, salahe = db.Exec("insert into transaksi values (?, ?, ?, ?)", no_nota, jenis, nama, bayar)
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Query Insert: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
   
    return response{
        Status: true,
        Pesan:  "Berhasil Tambah",
        Data:   []transaksi{},
    }
}
 
//fungsi ubah data
func ubah(no_nota string, jenis string, nama string, bayar string) response {
    db, salahe := koneksi()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Koneksi: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    defer db.Close()
 
    _, salahe = db.Exec("update transaksi set nama=?, jenis=?, bayar=? where no_nota=?", nama, jenis, bayar, no_nota)
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Query Update: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    return response{
        Status: true,
        Pesan:  "Berhasil Ubah",
        Data:   []transaksi{},
    }
}
 
//fungsi hapus data
func hapus(no_nota string) response {
    db, salahe := koneksi()
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Koneksi: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    defer db.Close()
    _, salahe = db.Exec("delete from transaksi where no_nota=?", no_nota)
    if salahe != nil {
        return response{
            Status: false,
            Pesan:  "Gagal Query Delete: " + salahe.Error(),
            Data:   []transaksi{},
        }
    }
 
    return response{
        Status: true,
        Pesan:  "Berhasil Hapus",
        Data:   []transaksi{},
    }
}
 
func kontroler(w http.ResponseWriter, r *http.Request) {
    var tampilHTML, salaheTampil = template.ParseFiles("template/tampil.html")
    if salaheTampil != nil {
        fmt.Println(salaheTampil.Error())
        return
    }
 
    var tambahHTML, salaheTambah = template.ParseFiles("template/tambah.html")
    if salaheTambah != nil {
        fmt.Println(salaheTambah.Error())
        return
    }
 
    var ubahHTML, salaheUbah = template.ParseFiles("template/ubah.html")
    if salaheUbah != nil {
        fmt.Println(salaheUbah.Error())
        return
    }
 
    var hapusHTML, salaheHapus = template.ParseFiles("template/hapus.html")
    if salaheHapus != nil {
        fmt.Println(salaheHapus.Error())
        return
    }
 
    switch r.Method {
        case "GET":aksi := r.URL.Query()["aksi"]
        if len(aksi) == 0 {
            tampilHTML.Execute(w, tampil("Berhasil Tampil"))
        } else if aksi[0] == "tambah" {
            tambahHTML.Execute(w, nil)
        } else if aksi[0] == "ubah" {
            id := r.URL.Query()["id"]
            ubahHTML.Execute(w, getSup(id[0]))
        } else if aksi[0] == "hapus" {
            id := r.URL.Query()["id"]
            hapusHTML.Execute(w, getSup(id[0]))
        } else {
            tampilHTML.Execute(w, tampil("Berhasil Tampil"))
        }
 
        case "POST":var salahe = r.ParseForm()
        if salahe != nil {
            fmt.Fprintln(w, "Kesalahan: ", salahe)
            return
        }
 
        var no_nota = r.FormValue("no_nota")
        var jenis   = r.FormValue("jenis")
        var nama    = r.FormValue("nama")
        var bayar   = r.FormValue("bayar")
        var aksi     = r.URL.Path
       
        if aksi == "/tambah" {
            var hasil = tambah(no_nota, jenis, nama, bayar)
            tampilHTML.Execute(w, tampil(hasil.Pesan))
        } else if aksi == "/ubah" {
            var hasil = ubah(no_nota, jenis, nama, bayar)
            tampilHTML.Execute(w, tampil(hasil.Pesan))
        } else if aksi == "/hapus" {
            var hasil = hapus(id)tampilHTML.Execute(w, tampil(hasil.Pesan))
        } else {
            tampilHTML.Execute(w, tampil("Berhasil Tampil"))
        }
 
        default:
            fmt.Fprint(w, "Maaf. Method yang didukung hanya GET dan POST")
    }
}
 
func main() {
    http.HandleFunc("/", kontroler)
    fmt.Println("Server berjalan di Port 8080...")
    http.ListenAndServe(":8080", nil)
}