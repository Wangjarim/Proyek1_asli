// Ambil nilai parameter dari URL
const urlParams = new URLSearchParams(window.location.search);
const Id = urlParams.get("id");

// Fetch data menggunakan ID dari URL parameter
fetch(`http://127.0.0.1:3000/api/belajar/${Id}`)
  .then((response) => response.json())
  .then((result) => {
    // console.log(result);
    // Mengatur nilai-nilai ke elemen-elemen di halaman "print_form.html"
    document.getElementById("nama_lengkap").textContent = result.data.nama_lengkap;
    document.getElementById("nik").textContent = result.data.nik;
    document.getElementById("jenis_kelamin").textContent = result.data.jenis_kelamin;
    document.getElementById("tempat_lahir").textContent = result.data.tempat_lahir;
    document.getElementById("tanggal_lahir").textContent = result.data.tanggal_lahir;
    document.getElementById("alamat").textContent = result.data.alamat;
    document.getElementById("no_hp").textContent = result.data.no_hp;
    document.getElementById("nama_dokter").textContent = result.data.jadwal_dokter.dokter.nama;
    document.getElementById("hari").textContent = result.data.jadwal_dokter.hari.hari;
    document.getElementById("jam").textContent = result.data.jadwal_dokter.jam.jam;
    document.getElementById("ruangan").textContent = result.data.jadwal_dokter.ruangan.nama_ruangan;
    document.getElementById("tgl_reservasi").textContent = result.data.tgl_reservasi;
  })
  .catch((error) => {
    console.error("Error:", error);
    alert("Terjadi kesalahan saat mengambil data.");
  });