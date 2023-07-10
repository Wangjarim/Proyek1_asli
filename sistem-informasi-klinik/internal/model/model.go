package model

import "time"

type Pasien struct {
	Id           int          `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Namalengkap  string       `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	Nik          int          `gorm:"column:nik" json:"nik"`
	Jeninkelamin string       `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Tempatlahir  string       `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	Tanggallahir string       `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	Alamat       string       `gorm:"column:alamat" json:"alamat"`
	Nohp         string       `gorm:"column:no_hp" json:"no_hp"`
	IdJadwal     string       `gorm:"column:id_jadwal;foreignKey:JadwalId" json:"id_jadwal"`
	TglReservasi time.Time    `gorm:"column:tgl_reservasi;default:CURRENT_TIMESTAMP" json:"tgl_reservasi"`
	JadwalDokter JadwalDokter `gorm:"foreignKey:IdJadwal" json:"jadwal_dokter"`
}

type Dokter struct {
	Id       int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Nid      int    `gorm:"column:nid" json:"nid"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Keahlian string `gorm:"column:keahlian" json:"keahlian"`
	Nohp     string `gorm:"column:no_hp" json:"no_hp"`
}

type Hari struct {
	Id   int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Hari string `gorm:"column:hari" json:"hari"`
}

type Jam struct {
	Id  int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Jam string `gorm:"column:jam" json:"jam"`
}

type Ruangan struct {
	KodeRuangan int    `gorm:"primaryKey;column:kode_ruangan;autoIncrement" json:"kode_ruangan"`
	NamaRuangan string `gorm:"column:nama_ruangan" json:"nama_ruangan"`
}

type JadwalDokter struct {
	Id        int     `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	IdDokter  int     `gorm:"column:id_dokter;foreignKey:DokterId" json:"id_dokter"`
	IdHari    int     `gorm:"column:id_hari;foreignKey:HariId" json:"id_hari"`
	IdJam     int     `gorm:"column:id_jam;foreignKey:JamId" json:"id_jam"`
	IdRuangan int     `gorm:"column:id_ruangan;foreignKey:RuanganId" json:"id_ruangan"`
	Dokter    Dokter  `gorm:"foreignKey:IdDokter" json:"dokter"`
	Hari      Hari    `gorm:"foreignKey:IdHari" json:"hari"`
	Jam       Jam     `gorm:"foreignKey:IdJam" json:"jam"`
	Ruangan   Ruangan `gorm:"foreignKey:IdRuangan" json:"ruangan"`
}
