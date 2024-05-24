<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Edit Parhalado</title>

    <!-- Custom fonts for this template-->
    <link href="../vendor/fontawesome-free/css/all.min.css" rel="stylesheet" type="text/css">
    <link
        href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i"
        rel="stylesheet">

    <!-- Custom styles for this template-->
    <link href="../css/sb-admin-2.min.css" rel="stylesheet">

</head>

<body id="page-top">

    <!-- Page Wrapper -->
    <div id="wrapper">

        <!-- Sidebar -->
        <ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">
            <!-- Sidebar - Brand -->
            <a class="sidebar-brand d-flex align-items-center justify-content-center" href="index.html">
                <div class="logo">
                    <img src="../img/LOGO_HKBP.png" alt="Logo" style="max-width: 63px; max-height: 63px" />
                </div>
                <div class="sidebar-brand-text mx-3">HKBP Parparean</div>
            </a>

            <!-- Divider -->
            <hr class="sidebar-divider my-0" />


            <li class="nav-item ">
                <a class="nav-link" href="/berita">
                    <i class="fas fa-newspaper"></i>
                    <span>Berita</span></a>
            </li>

            <li class="nav-item ">
                <a class="nav-link" href="/acara">
                    <i class="fas fa-calendar"></i>
                    <span>Acara</span></a>
            </li>


            <li class="nav-item ">
                <a class="nav-link" href="/pengumuman">
                    <i class="fas fa-bullhorn"></i>
                    <span>Pengumuman</span></a>
            </li>

            <li class="nav-item ">
                <a class="nav-link" href="/renungan">
                    <i class="fas fa-book"></i>
                    <span>Renungan</span></a>
            </li>

            <li class="nav-item ">
                <a class="nav-link" href="/sektor">
                    <i class="fas fa-map-marker"></i>
                    <span>Sektor</span></a>
            </li>

            <li class="nav-item active">
                <a class="nav-link" href="/parhalado">
                    <i class="fas fa-users"></i>
                    <span>Parhalado</span></a>
            </li>


            <!-- Divider -->
            <hr class="sidebar-divider d-none d-md-block" />

            <!-- Sidebar Toggler (Sidebar) -->
            <div class="text-center d-none d-md-inline">
                <button class="rounded-circle border-0" id="sidebarToggle"></button>
            </div>
        </ul>
        <!-- End of Sidebar -->

        <!-- Content Wrapper -->
        <div id="content-wrapper" class="d-flex flex-column">
            <!-- Topbar -->
            <nav class="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow">
                <!-- Topbar Navbar -->
                <ul class="navbar-nav ml-auto">

                    <div class="topbar-divider d-none d-sm-block"></div>
                    <!-- Nav Item - User Information -->
                    <li class="nav-item dropdown no-arrow">
                        <a class="nav-link dropdown-toggle" href="#" id="userDropdown" role="button"
                            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            <span class="mr-2 d-none d-lg-inline text-gray-600 small">User Admin</span>
                                <img class="img-profile rounded-circle" src="../img/logo-hkbp.png">
                        </a>
                    </li>
                </ul>
            </nav>
            <!-- End of Topbar -->
            <main>
                <div class="container-fluid">
                    <nav aria-label="breadcrumb" class="breadcrumb-box d-flex justify-content-lg-end">
                        <ol class="breadcrumb">
                            <li class="breadcrumb-item">
                                    <a href=../acara>Daftar Parhalado</a> </li> <li class="breadcrumb-item active"
                                        aria-current="page">
                                        <a>Edit Parhalado</a>
                            </li>
                        </ol>
                    </nav>
                    <!-- Basic Card Example -->
                    <div class="card shadow col-xl-10 col-md-6 mb-4">
                        <div class="card-header py-3">
                            <h6 class="m-0 font-weight-bold text-primary">Edit Daftar Parhalado</h6>
                        </div>
                        <div class="card-body">
                            <form action="{{ route('edit.parhalado', $data['ID']) }}" method="POST"
                                enctype="multipart/form-data">
                                @csrf
                                @method('PUT')
                                <div class="form-group row">
                                    <label for="nama" class="col-sm-2 col-form-label">Nama</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control" id="nama" name="nama"
                                            value="{{ $data['nama'] }}" required>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label for="tanggal_lahir" class="col-sm-2 col-form-label">Tanggal Lahir</label>
                                    <div class="col-sm-10">
                                        <input type="date" class="form-control " id="tanggal_lahir" name="tanggal_lahir"
                                            value="{{ $data['tanggal_lahir'] }}" required>
                                    </div>
                                </div>
                                <div class="form-group row">

                                    <label for="jabatan" class="col-sm-2 col-form-label">Jabatan</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control " id="jabatan" name="jabatan"
                                            value="{{ $data['jabatan'] }}" required>
                                    </div>
                                </div>
                                <div class="form-group row">

                                    <label for="alamat" class="col-sm-2 col-form-label">Alamat</label>
                                    <div class="col-sm-10">
                                        <input type="text" class="form-control " id="alamat" name="alamat"
                                            value="{{ $data['alamat'] }}" required>
                                    </div>
                                </div>

                        </div>
                    </div>



                    <div class="form-group row ">
                        <div class="col-sm-10">
                            <button type="submit" name="submit" id="submit" class="btn btn-success">Simpan
                                Perubahan</button>
                        </div>
                    </div>
                    </form>
                </div>
        </div>
    </div>
    </main>

    <!-- Footer -->
    <footer class="sticky-footer bg-white">
        <div class="container my-auto">
            <div class="copyright text-center my-auto">
                <span>Copyright &copy; Your Website 2020</span>
            </div>
        </div>
    </footer>
    <!-- End of Footer -->

    </div>
    <!-- End of Content Wrapper -->

    </div>
    <!-- End of Page Wrapper -->

    <!-- Scroll to Top Button-->
    <a class="scroll-to-top rounded" href="#page-top">
        <i class="fas fa-angle-up"></i>
    </a>

    <!-- Logout Modal-->
    <div class="modal fade" id="logoutModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel"
        aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLabel">Ready to Leave?</h5>
                    <button class="close" type="button" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">×</span>
                    </button>
                </div>
                <div class="modal-body">Select "Logout" below if you are ready to end your current session.</div>
                <div class="modal-footer">
                    <button class="btn btn-secondary" type="button" data-dismiss="modal">Cancel</button>
                    <a class="btn btn-primary" href="login.html">Logout</a>
                </div>
            </div>
        </div>
    </div>

    <!-- Bootstrap core JavaScript-->
    <script src="../vendor/jquery/jquery.min.js"></script>
    <script src="../vendor/bootstrap/js/bootstrap.bundle.min.js"></script>

    <!-- Core plugin JavaScript-->
    <script src="../vendor/jquery-easing/jquery.easing.min.js"></script>

    <!-- Custom scripts for all pages-->
    <script src="../js/sb-admin-2.min.js"></script>

</body>

</html>
