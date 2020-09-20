$(document).ready(function () {
    if ($('#btCustomers').length) {
        $('#btCustomers').DataTable({
            bFilter: false,
            aLengthMenu: [[10, 25, 50, -1], [10, 25, 50, "All"]],
        });
    }

    if ($('#birth_date').length) {
        $('#birth_date').datepicker({
            orientation: "bottom auto",
            format: "yyyy-mm-dd",
        });
    }
});