
async function createNote() {
    var title = document.querySelector("#input-title").value;
    var element = document.querySelector(".trix-custom-content");
    // var document = element.editor.getDocument();
    console.log(title);
    console.log(element);
    // const data = {
    //     title: title,
    //     body: content,
    // }
    // await axios.post(ROOT_API + "api/notes/", data)
    // .then(response => {
    //     if (response.status == 201) {
    //         window.location.href = ROOT_API + "/view-notes";
    //     }
    // });
}