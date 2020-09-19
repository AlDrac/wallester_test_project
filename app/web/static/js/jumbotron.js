$(document).ready(function () {
    if ($('#btCustomers').length) {
        $('#btCustomers').DataTable({
            aLengthMenu: [[10, 25, 50, -1], [10, 25, 50, "All"]]
        });
    }
});