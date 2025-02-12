document.addEventListener("DOMContentLoaded", function (){
    const toggle = document.getElementById("toggle")
    const close = document.getElementById("open-toggle")
    const open = document.getElementById("close-toggle")
    const menu = document.getElementById("mobile-menu")
    const search = document.getElementById("search")

    const homelink = document.getElementById("home-link");
    const bloglink = document.getElementById("blog-link");
    const projectslink = document.getElementById("projects-link");
    const contactlink = document.getElementById("contact-link");

// Function to remove all active classes
    function removeActiveClasses() {
        homelink.classList.remove("border-indigo-500", "text-indigo-500");
        bloglink.classList.remove("border-indigo-500", "text-indigo-500");
        projectslink.classList.remove("border-indigo-500", "text-indigo-500");
        contactlink.classList.remove("border-indigo-500", "text-indigo-500");

        homelink.classList.add("border-transparent");
        bloglink.classList.add("border-transparent");
        projectslink.classList.add("border-transparent");
        contactlink.classList.add("border-transparent");
    }

// Check the current path and apply the active class to the corresponding link
    if (window.location.pathname === "/") {
        removeActiveClasses();
        homelink.classList.add("border-indigo-500", "text-indigo-500");
        homelink.classList.remove("border-transparent");
    } else if (window.location.pathname === "/blog") {
        removeActiveClasses();
        bloglink.classList.add("border-indigo-500", "text-indigo-500");
        bloglink.classList.remove("border-transparent");
    } else if (window.location.pathname === "/projects") {
        removeActiveClasses();
        projectslink.classList.add("border-indigo-500", "text-indigo-500");
        projectslink.classList.remove("border-transparent");
    } else if (window.location.pathname === "/contact") {
        removeActiveClasses();
        contactlink.classList.add("border-indigo-500", "text-indigo-500");
        contactlink.classList.remove("border-transparent");
    }

    toggle.addEventListener("click", function (){
        close.classList.toggle("hidden");
        open.classList.toggle("hidden");
        menu.classList.toggle("hidden");
    })

    search.addEventListener("keydown", function (event) {
        if (event.code !== "Enter") return;

        window.location.href = "https://robertgee.dev/projects?search=" + search.value
    })
})