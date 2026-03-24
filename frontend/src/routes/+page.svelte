<script>
	import { onMount } from 'svelte';

	let users = $state([]);
	let loading = $state(false);
	let name = $state("");
	let error = $state("");

	let editingId = $state(null);
	let editName = $state("");

	async function getUsers() {
		loading = true;
		error = "";

		try {
			const res = await fetch('http://localhost:3000/users');
			if (!res.ok) throw new Error("Gagal fetch data");

			users = await res.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	async function addUser() {
		if (!name.trim()) {
			error = "Nama tidak boleh kosong";
			return;
		}

		error = "";

		try {
			const res = await fetch('http://localhost:3000/users', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ name })
			});

			if (!res.ok) throw new Error("Gagal menambah user");

			name = "";
			await getUsers();
		} catch (err) {
			error = err.message;
		}
	}

	async function deleteUser(id) {
		try {
			const res = await fetch(`http://localhost:3000/users/${id}`, {
				method: 'DELETE'
			});

			if (!res.ok) throw new Error("Gagal hapus");

			await getUsers();
		} catch (err) {
			error = err.message;
		}
	}

	function startEdit(user) {
		editingId = user.id;
		editName = user.name;
	}

	function cancelEdit() {
		editingId = null;
		editName = "";
	}

	async function updateUser(id) {
		if (!editName.trim()) {
			error = "Nama tidak boleh kosong";
			return;
		}

		try {
			const res = await fetch(`http://localhost:3000/users/${id}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ name: editName })
			});

			if (!res.ok) throw new Error("Gagal update");

			editingId = null;
			editName = "";
			await getUsers();
		} catch (err) {
			error = err.message;
		}
	}

	onMount(getUsers);
</script>

<style>
	.container {
		max-width: 600px;
		margin: auto;
		padding: 20px;
		font-family: Arial, sans-serif;
	}

	h1 {
		text-align: center;
	}

	.card {
		background: #fff;
		padding: 15px;
		border-radius: 10px;
		box-shadow: 0 4px 10px rgba(0,0,0,0.1);
		margin-top: 20px;
	}

	input {
		padding: 10px;
		border-radius: 8px;
		border: 1px solid #ccc;
	}

	.input-main {
		width: 70%;
		margin-right: 10px;
	}

	button {
		padding: 8px 12px;
		border: none;
		border-radius: 8px;
		cursor: pointer;
	}

	.btn-add {
		background: #4CAF50;
		color: white;
	}

	.btn-delete {
		background: #ff4d4d;
		color: white;
		margin-left: 5px;
	}

	.btn-edit {
		background: #2196F3;
		color: white;
		margin-left: 5px;
	}

	.btn-save {
		background: #4CAF50;
		color: white;
		margin-left: 5px;
	}

	.btn-cancel {
		background: #aaa;
		color: white;
		margin-left: 5px;
	}

	.user-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 10px 0;
		border-bottom: 1px solid #eee;
	}

	.error {
		color: red;
		margin-top: 10px;
		text-align: center;
	}

	.loading {
		text-align: center;
		margin-top: 20px;
	}
</style>

<div class="container">
	<h1>👤 User Management</h1>

	<div class="card">
		<input class="input-main" bind:value={name} placeholder="Masukkan nama user..." />
		<button class="btn-add" onclick={addUser}>Tambah</button>

		{#if error}
			<p class="error">{error}</p>
		{/if}
	</div>

	<div class="card">
		{#if loading}
			<p class="loading">Loading...</p>
		{:else if users.length === 0}
			<p class="loading">Belum ada user</p>
		{:else}
			{#each users as user (user.id)}
				<div class="user-item">
					{#if editingId === user.id}
						<input bind:value={editName} />
						<div>
							<button class="btn-save" onclick={() => updateUser(user.id)}>Simpan</button>
							<button class="btn-cancel" onclick={cancelEdit}>Batal</button>
						</div>
					{:else}
						<span>{user.id} - {user.name}</span>
						<div>
							<button class="btn-edit" onclick={() => startEdit(user)}>Edit</button>
							<button class="btn-delete" onclick={() => deleteUser(user.id)}>Hapus</button>
						</div>
					{/if}
				</div>
			{/each}
		{/if}

		<p style="margin-top:10px;"><b>Total:</b> {users.length}</p>
	</div>
</div>