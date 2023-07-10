// Ambil elemen formulir
const form = document.getElementById("contact-form");
const iddokterInput = document.getElementById("id_dokter");
const idhariInput = document.getElementById("id_hari");
const idjamInput = document.getElementById("id_jam");
const idruanganInput = document.getElementById("id_ruangan");
// Tambahkan event listener pada tombol "Input Data"
const inputButton = document.querySelector(".btn-success");
inputButton.addEventListener("click", () => {
  // Ambil nilai input dari formulir
  const id_dokter = parseInt(iddokterInput.value);
  const id_hari = parseInt(idhariInput.value);
  const id_jam = parseInt(idjamInput.value);
  const id_ruangan = parseInt(idruanganInput.value);
  // Validasi input
  if (!id_dokter || !id_hari || !id_jam || !id_ruangan ) {
    alert("Silakan lengkapi semua field");
    return;
  }
  // Buat objek data yang akan dikirim ke server
  const data = {
    id_dokter: id_dokter,
    id_hari: id_hari,
    id_jam: id_jam,
    id_ruangan: id_ruangan,
  };
  // Kirim permintaan POST ke server untuk menambahkan data
  fetch("http://127.0.0.1:3000/api/jadwal_dokter", {
    method: "POST",
    headers: {
      "Content-Type": "application/json", 
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((result) => {
      console.log("Sukses:", result);
      alert("Data berhasil ditambahkan!");
      window.location.href = "index_jadwal.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat menambahkan data.");
    });
});