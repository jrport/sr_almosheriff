export const ssr = false
export async function load({ fetch, params }) {
	const res = await fetch("http://localhost:8080/pedidos");
	const pedidos = await res.json();

	return {pedidos};
}
