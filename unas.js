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
        console.log("Enviando datos del formulario:", formData);

        const res = await fetch('http://localhost:8080/api/citas', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(formData)
        });

        console.log("Respuesta recibida:", res);

        let data;
        const text = await res.text(); // leemos como texto para evitar el error
        if (text) {
            data = JSON.parse(text); // lo convertimos a JSON solo si hay contenido
        } else {
            data = { message: "Sin respuesta JSON del servidor" };
        }

        console.log("Respuesta procesada:", data);

    } catch (error) {
        console.error("Error enviando la reserva:", error);
    }
});