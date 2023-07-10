// Ambil elemen formulir
const form = document.getElementById("contact-form");
const nidInput = document.getElementById("nid");
const namaInput = document.getElementById("nama");
const keahlianInput = document.getElementById("keahlian");
const nohpInput = document.getElementById("no_hp");
// Tambahkan event listener pada tombol "Input Data"
const inputButton = document.querySelector(".btn-success");
inputButton.addEventListener("click", () => {
  // Ambil nilai input dari formulir
  const nid = parseInt(nidInput.value);
  const nama = namaInput.value;
  const keahlian = keahlianInput.value;
  const no_hp = nohpInput.value;
  // Validasi input
  if (!nid || !nama || !keahlian || !no_hp) {
    alert("Silakan lengkapi semua field");
    return;
  }
  // Buat objek data yang akan dikirim ke server
  const data = {
    nid: nid,
    nama: nama,
    keahlian: keahlian,
    no_hp: no_hp,
  };
  // Kirim permintaan POST ke server untuk menambahkan data
  fetch("http://127.0.0.1:3000/api/dokter", {
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
      window.location.href = "index_doktor.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Terjadi kesalahan saat menambahkan data.");
    });
});
