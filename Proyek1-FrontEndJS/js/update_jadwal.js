// Ambil data dari URL
const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("id");

// Ambil elemen formulir
const form = document.getElementById("contact-form");
const iddokterInput = document.getElementById("id_dokter");
const idhariInput = document.getElementById("id_hari");
const idjamInput = document.getElementById("id_jam");
const idruanganInput = document.getElementById("id_ruangan");
// Isi formulir dengan data dari server berdasarkan ID
fetch(`http://127.0.0.1:3000/api/jadwal_dokter/${id}`)
.then((result) => {
    return result.json();
  })
  .then((data) => {
    const jadwal_dokter = data.data;
    iddokterInput.value = jadwal_dokter.id_dokter; 
    idhariInput.value = jadwal_dokter.id_hari;
    idjamInput.value = jadwal_dokter.id_jam;
    idruanganInput.value = jadwal_dokter.id_ruangan;
  });
// Fungsi untuk mengirim data ke server
function updateData() {
  // Buat objek data yang akan dikirim ke server
  const data = {
    id_dokter: parseInt(iddokterInput.value),
    id_hari: idhariInput.value,
    id_jam: idjamInput.value,
    id_ruangan: idruanganInput.value,
  };
  // Buat konfigurasi untuk request
  const options = {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };
  // Kirim request ke server
  fetch(`http://127.0.0.1:3000/api/jadwal_dokter/${id}`, options)
    .then((result) => {
      return result.json();
    })
    .then((data) => {
      alert("Data berhasil diupdate");
      window.location.href = "index_jadwal.html";
    })
    .catch((error) => {
      console.error(error);
      alert("Data gagal diupdate");
    });
}
// Tambahkan event listener pada tombol Edit
const editButton = document.querySelector(".btn-success");
editButton.addEventListener("click", updateData);
