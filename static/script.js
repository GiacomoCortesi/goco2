async function fetchData() {
    const response = await fetch('/saving');
    const data = await response.json();
    document.getElementById('daySaving').textContent = `${data.day} Kg`;
    document.getElementById('weekSaving').textContent = `${data.week} Kg`;
    document.getElementById('monthSaving').textContent = `${data.month} Kg`;
    document.getElementById('yearSaving').textContent = `${data.year} Kg`;
}

async function addIntervention() {
    const date = document.getElementById('dateInput').value;
    const id = document.getElementById('idInput').value;

    const response = await fetch('/intervention', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            date: date ? date : null,
            id: id ? id: null,
        }),
    });

    if (response.ok) {
        fetchData()
        alert('Intervention added successfully!');
    } else {
        alert('Failed to add intervention. Please try again.');
    }
}

fetchData();
