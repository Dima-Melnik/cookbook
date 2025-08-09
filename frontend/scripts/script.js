const recipesList = document.getElementById('recipeList');
const recipeForm = document.getElementById('addRecipeForm');
const closeBtn = document.querySelector('.close-btn');
const titleInput = document.getElementById('titleInput');
const descInput = document.getElementById('descriptionInput');
const addBtn = document.getElementById('addButton');

const modalOverlay = document.getElementById('modalRecipeOverlay');
const modalTitle = document.getElementById('modalRecipeTitle');
const modalDesc = document.getElementById('modalRecipeDesc');
const closeModalBtn = document.getElementById('closeModalRecipe');

const API_URL = 'http://localhost:8080/';

window.addEventListener('DOMContentLoaded', () => {
    getAllRecipes();
});

async function getAllRecipes() {
    try {
        const response = await fetch(API_URL, { method: 'GET' });

        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }

        const recipes = await response.json();

        console.log("SERVER RESPONSE:", recipes);

        if (recipes.type !== "success") {
            throw new Error("Server returned error");
        }

        renderRecipesList(recipes.message);
    } catch (error) {
        console.error('Error loading recipes list:', error);
        recipesList.innerHTML = `<li style="color:red;">Failed to load recipes list</li>`;
    }
}

function renderRecipesList(recipes) {
    recipesList.innerHTML = '';

    recipes.forEach(recipe => {
        const li = document.createElement('li');

        li.textContent = recipe.title;
        li.style.cursor = 'pointer';
        li.addEventListener('click', () => {
            openRecipeModal(recipe.title, recipe.description);
        });
        closeModalBtn.addEventListener('click', () => {
            closeRecipeModal();
        })

        const deleteBtn = document.createElement('button');
        deleteBtn.textContent = 'Delete';
        deleteBtn.style.marginLeft = '50px';
        deleteBtn.addEventListener('click', () => {
            deleteRecipe(recipe.id);
        });

        li.appendChild(deleteBtn);
        recipesList.appendChild(li);
    });
}

// Відкриття модалки
function openRecipeModal(title, desc) {
    modalTitle.textContent = title;
    modalDesc.textContent = desc;
    modalOverlay.style.display = 'flex';
}

// Закриття модалки
function closeRecipeModal() {
    modalOverlay.style.display = 'none';
}

recipeForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const title = titleInput.value.trim();
    const description = descInput.value.trim();

    if (!title || !description) {
        alert('Please fill in both fields');
        return;
    }

    try {
        const response = await fetch(API_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, description }),
        });

        if (!response.ok) throw new Error(`Error: ${response.status}`);

        const result = await response.json();
        if (result.type !== "success") {
            throw new Error(result.message || "Server returned error");
        }

        titleInput.value = '';
        descInput.value = '';
        recipeForm.classList.add('hidden');

        getAllRecipes();
    } catch (error) {
        console.error('Error creating recipe', error);
        alert('Failed creating recipe');
    }
});

async function deleteRecipe(id) {
    if (!confirm('Are you sure you want to delete this recipe?')) return;

    try {
        const response = await fetch(`${API_URL}${id}`, {
            method: 'DELETE',
        });

        if (!response.ok) throw new Error(`Error: ${response.status}`);

        const result = await response.json();
        if (result.type !== "success") {
            throw new Erroro(result.message || "Server retunred error");
        } 

        getAllRecipes();
    } catch (error) {
        console.error('Error deleting recipe:', error);
        alert('Failed deleting recipe');
    }
}

addBtn.addEventListener('click', () => {
    recipeForm.classList.toggle('hidden');
});

closeBtn.addEventListener('click', () => {
    recipeForm.classList.add('hidden');
});
