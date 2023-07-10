// Ambil data dari URL
const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("id");

// Ambil elemen formulir
const form = document.getElementById("contact-form");
const nidInput = document.getElementById("nid");
const namaInput = document.getElementById("nama");
const keahlianInput = document.getElementById("keahlian");
const nohpInput = document.getElementById("no_hp");
// Isi formulir dengan data dari server berdasarkan ID
fetch(`http://127.0.0.1:3000/api/dokter/${id}`)
.then((result) => {
    return result.json();
  })
  .then((data) => {
    const dokter = data.data;
    nidInput.value = dokter.nid; 
    namaInput.value = dokter.nama;
    keahlianInput.value = dokter.keahlian;
    nohpInput.value = dokter.no_hp;
  });
// Fungsi untuk mengirim data ke server
function updateData() {
  // Buat objek data yang akan dikirim ke server
  const data = {
    nid: parseInt(nidInput.value),
    nama: namaInput.value,
    keahlian: keahlianInput.value,
    no_hp: nohpInput.value,
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
  fetch(`http://127.0.0.1:3000/api/dokter/${id}`, options)
    .then((result) => {
      return result.json();
    })
    .then((data) => {
      alert("Data berhasil diupdate");
      window.location.href = "index_doktor.html";
    })
    .catch((error) => {
      console.error(error);
      alert("Data gagal diupdate");
    });
}
// Tambahkan event listener pada tombol Edit
const editButton = document.querySelector(".btn-success");
editButton.addEventListener("click", updateData);
