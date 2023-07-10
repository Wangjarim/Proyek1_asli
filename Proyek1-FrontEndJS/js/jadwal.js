

fetch("http://127.0.0.1:3000/api/jadwal_dokter")
	.then((result) => {
		return result.json();
	})
	.then((data) => {
		let tableData = "";
		data.data.map((values) => {
			// Ambil informasi nama, hari, dan jam dari respons JSON
			const namaDokter = values.dokter.nama;
			const namaHari = values.hari.hari;
			const jam = values.jam.jam;
			const ruangan = values.ruangan.nama_ruangan;

			// Manipulasi data pegawai dan masukkan ke dalam bentuk tabel
			tableData += `
                <tr>
                <td>${namaDokter}</td>
                    <td>${namaHari}</td>
                    <td>${jam}</td>
                    <td>${ruangan}</td>
				<td>
				<button type="button" class="btn btn-success" jadwal-id="${values.id}">Reservasi</button>
				</td>

                </tr>`;
		});
		// Tampilkan data pegawai ke dalam tabel
		document.getElementById("jadwal-table").innerHTML = tableData;

		// Tambahkan event listener pada setiap tombol delete
		const deleteButtons = document.querySelectorAll('.btn-danger');
		deleteButtons.forEach((button) => {
			button.addEventListener('click', (event) => {
				const id = event.target.getAttribute('jadwal_dokter-id');
				// Kirim permintaan DELETE ke server
				fetch(`http://127.0.0.1:3000/api/jadwal_dokter/${id}`, {
					method: 'DELETE'
				})
					.then((response) => {
						if (response.ok) {
							// Sukses menghapus data, lakukan tindakan sesuai kebutuhan (misalnya, memperbarui tampilan tabel)
							// Contoh:
							event.target.parentNode.parentNode.remove(); // Menghapus baris tabel yang sesuai dengan tombol delete yang diklik
						} else {
							// Gagal menghapus data, berikan umpan balik kepada pengguna atau tangani kesalahan
							throw new Error('Gagal menghapus data');
						}
					})
					.catch(error => {
						console.log('error', error);
						alert('Terjadi kesalahan pada server');
					});
			});
		});

		// Tambahkan event listener pada setiap tombol detail
		const detailButtons = document.querySelectorAll('.btn-success');
		detailButtons.forEach((button) => {
			button.addEventListener('click', (event) => {
				const id = event.target.getAttribute('jadwal-id');
				window.location.href = `input.html?id=${id}`;
			});
		});
	})
	.catch(error => {
		console.log('error', error);
		alert('Terjadi kesalahan pada server');
	});