const addForm = document.getElementById("add-form")
const form = document.getElementById("form")
let isAdding = false

addForm.addEventListener("click", () => {
	if(isAdding) return

	isAdding = true

	const formElement = document.createElement("form")
	formElement.classList.add("flex")

	const inputForm = document.createElement("input")
	inputForm.classList.add("flex-1", "outline-none", "bg-transparent", "border-b", "border-black")
	inputForm.setAttribute("name", "description")

	const addButton = document.createElement("button")
	addButton.classList.add("select-none", "py-2", "px-4", "rounded", "border", "border-green-300", "bg-green-300", "text-white", "transition-all", "hover:bg-green-400", "hover:border-green-400", "mx-3")
	addButton.type = "submit"
	addButton.textContent = "Add"
	addButton.addEventListener("click", () => {
		// make request to /add endpoint
	})

	const cancelButton = document.createElement("button")
	cancelButton.classList.add("select-none", "py-2", "px-4", "rounded", "border", "border-red-300", "bg-red-300", "text-white", "transition-all", "hover:bg-red-400", "hover:border-red-400")
	cancelButton.type = "submit"
	cancelButton.textContent = "Cancel"
	cancelButton.addEventListener("click", () => {
		form.removeChild(formElement)
		isAdding = false
	})

	formElement.appendChild(inputForm)
	formElement.appendChild(addButton)
	formElement.appendChild(cancelButton)
	form.appendChild(formElement)
})