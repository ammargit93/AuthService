const loginForm = document.getElementById('loginForm');

loginForm.addEventListener('submit', async (event) => {
    event.preventDefault();  // Prevent the default form submission

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    console.log(username)
    console.log(password)
    // Make a POST request to the login endpoint of your Go service
    const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),  // Send the credentials
    });

    const data = await response.json();  // Get the response data
    console.log(data)

    if (response.ok) {
        // Store the token in localStorage
        localStorage.setItem('authToken', data.token);
        alert('Login successful!');
    } else {
        // If the login fails, show an error message
        alert('Login failed: ' + data.error);
    }
});

// Function to automatically attach token to requests
async function apiFetch(url, options = {}) {
    const token = localStorage.getItem('authToken');  // Get the token from localStorage

    if (token) {
        options.headers = {
            ...options.headers,
            'Authorization': `Bearer ${token}`,  // Attach the token in the Authorization header
        };
    }

    const response = await fetch(url, options);  // Make the API call
    return response;
}

// Example of using apiFetch to call a protected route
async function fetchProtectedData() {
    const response = await apiFetch('http://localhost:8080/protected', {
        method: 'GET',  // Use GET for fetching protected data
    });

    if (response.ok) {
        const data = await response.json();
        console.log('Protected data:', data);  // Log protected data
    } else {
        alert('Access denied or token expired');
    }
}
