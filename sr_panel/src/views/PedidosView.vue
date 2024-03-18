<template>
<div class="text-black w-full">
	<v-frappe-chart
		type="bar"
		:labels="['Fraga Maia', 'Avenida', 'Senhor dos P.', 'Seabra', 'Anel de Cont.', 'Sim', 'Mochila', 'Feira X']"
		:data="[{ values: [0,0,0,0,0,0,0,0] }]"
		:colors="['red']"
		v-if="loading"
	/>
	<v-frappe-chart
		type="bar"
		:labels="['Fraga Maia', 'Avenida', 'Senhor dos P.', 'Seabra', 'Anel de Cont.', 'Sim', 'Mochila', 'Feira X']"
		:data="[{{ pedidos }}]"
		:colors="['red']"
		v-if="pedidos"
	/>
	<div v-if="error" class="error">{{ error }}</div>
</div>
</template>

<script setup>
	import { ref, watch } from 'vue'

	const loading = ref(false)
	const pedidos = ref(null)
	const error = ref(null)	

	async function fetchData(id) {
		error.value = pedidos.value = null
		loading.value = true
		
		try {
			// replace `getPost` with your data fetching util / API wrapper
			 pedidos.value = await fetch('localhost:8080/pedidos')
		} catch (err) {
			error.value = err.toString()
		} finally {
			loading.value = false
		}
	}

	fetchData()
</script>
