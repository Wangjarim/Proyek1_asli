// const urlParams = new URLSearchParams(window.location.search);
// const id = urlParams.get("id");

// Ambil elemen formulir
const form = document.getElementById("contact-form");
const namalengkapInput = document.getElementById("nama_lengkap");
const nikInput = document.getElementById("nik");
const jeniskelaminInput = document.getElementById("jenis_kelamin");
const tempatlahirInput = document.getElementById("tempat_lahir");
const tanggallahirInput = document.getElementById("tanggal_lahir");
const alamatInput = document.getElementById("alamat");
const nohpInput = document.getElementById("no_hp");
const idjadwalInput = document.getElementById("id_jadwal");
const nama_dokter = document.getElementById("nama_dokter");
const hari = document.getElementById("hari");
const jam = document.getElementById("jam");
const ruangan = document.getElementById("ruangan");
// const tglreservasiInput = document.getElementById("tgl_reservasi");
// Tambahkan event listener pada tombol "Input Data"
const inputButton = document.querySelector(".btn-success");

let queryString = {}
location.search.substr(1).split("&").forEach(function(item) {queryString[item.split("=")[0]] = item.split("=")[1]})

fetch("http://127.0.0.1:3000/api/jadwal_dokter/"+ queryString.id)
  .then((result) => {
    return result.json();
  })
  .then(({data}) => {
    // console.log(data)
    idjadwalInput.value = data.id;
    nama_dokter.value = data.dokter.nama;
    hari.value = data.hari.hari;
    jam.value = data.jam.jam;
    ruangan.value = data.ruangan.nama_ruangan;
  });

inputButton.addEventListener("click", () => { 
  // Ambil nilai input dari formulir
  const nama_lengkap = namalengkapInput.value;
  const nik = parseInt(nikInput.value);
  const jenis_kelamin = jeniskelaminInput.value;
  const tempat_lahir = tempatlahirInput.value;
  const tanggal_lahir = tanggallahirInput.value;
  const alamat = alamatInput.value;
  const no_hp = nohpInput.value;
  const id_jadwal = idjadwalInput.value;
  // const tgl_reservasi = tglreservasiInput.value;
  // Validasi input
  if (
    !nama_lengkap ||
    !nik ||
    !jenis_kelamin ||
    !tempat_lahir ||
    !tanggal_lahir ||
    !alamat ||
    !no_hp ||
    !id_jadwal
    // !tgl_reservasi
  ) {
    alert("Silakan lengkapi semua field");
    return;
  }
  // Buat objek data yang akan dikirim ke server
  const data = {
    nama_lengkap: nama_lengkap,
    nik: nik,
    jenis_kelamin: jenis_kelamin,
    tempat_lahir: tempat_lahir,
    tanggal_lahir: tanggal_lahir,
    alamat: alamat,
    no_hp: no_hp,
    id_jadwal: id_jadwal,
    // tgl_reservasi: tgl_reservasi,
  };
//cobain developinp
  // Kirim permintaan POST ke server untuk menambahkan data
  fetch("http://127.0.0.1:3000/api/belajar/", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((result) => {
      console.log("Sukses:", result);
      const id_passien = result.data
      window.location.href = `print_form.html?id=${id_passien}`;
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat menambahkan data.");
    });
});