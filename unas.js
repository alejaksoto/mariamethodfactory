document.getElementById('formReserva').addEventListener('submit', async (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = {
        nombre: form.nombre.value,
        servicio: form.servicio.value,
        fecha: form.fecha.value,
        telefono: form.telefono.value
    };

    try {
        const res = await fetch('http://localhost:8080/api/citas', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(formData)
        });

        const result = await res.json();
        alert(result.message || "Reserva enviada con éxito");
        form.reset();
    } catch (error) {
        console.error("Error enviando la reserva:", error);
        alert("No se pudo enviar la reserva, intenta de nuevo.");
    }
});
