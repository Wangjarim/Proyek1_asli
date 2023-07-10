

fetch("http://127.0.0.1:3000/api/belajar")
	.then((result) => {
		return result.json();
	})
	.then((data) => {
		let tableData = "";
		data.data.map((values) => {
			// Manipulasi data pegawai dan masukkan ke dalam bentuk tabel
			tableData += `
                <tr>
                <td>${values.nama_lengkap}</td>
                <td>${values.nik}</td>
                <td>${values.jenis_kelamin}</td>
                <td>${values.tempat_lahir}</td>
                <td>${values.tanggal_lahir}</td>
                <td>${values.alamat}</td>
                <td>${values.no_hp}</td>
                <td>${values.jadwal_dokter.dokter.nama}</td>
                <td>${values.jadwal_dokter.hari.hari}</td>
                <td>${values.jadwal_dokter.jam.jam}</td>
                <td>${values.jadwal_dokter.ruangan.nama_ruangan}</td>
                <td>${values.tgl_reservasi}</td>
                <td>
                <button type="button" class="btn btn-danger" belajar-id="${values.id}">Delete</button>
            </td>
                </tr>`;
		});
		// Tampilkan data pegawai ke dalam tabel
		document.getElementById("index-table").innerHTML = tableData;

		// Tambahkan event listener pada setiap tombol delete
		const deleteButtons = document.querySelectorAll('.btn-danger');
		deleteButtons.forEach((button) => {
			button.addEventListener('click', (event) => {
				const id = event.target.getAttribute('belajar-id');
				// Kirim permintaan DELETE ke server
				fetch(`http://127.0.0.1:3000/api/belajar/${id}`, {
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
				const id_user = event.target.getAttribute('belajar-id');
				window.location.href = `update.html?id=${id_user}`;
			});
		});
	})
	.catch(error => {
		console.log('error', error);
		alert('Terjadi kesalahan pada server');
	});