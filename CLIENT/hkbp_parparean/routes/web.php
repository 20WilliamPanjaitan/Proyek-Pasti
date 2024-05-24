<?php

use App\Http\Controllers\AcaraController;
use App\Http\Controllers\BeritaController;
use App\Http\Controllers\ParhaladoController;
use App\Http\Controllers\PengumumanController;
use App\Http\Controllers\RenunganController;
use App\Http\Controllers\SektorController;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/



// BERITA
Route::get('/', [BeritaController::class, 'getBeritaApi']);
Route::get('/berita', [BeritaController::class, 'getBeritaApi'])->name('berita');
Route::delete('/berita/{id}', [BeritaController::class, 'deleteBerita'])->name('hapus.berita');
Route::get('/berita/{id}', [BeritaController::class, 'showUpdate'])->name('show.berita');
Route::put('/editberita/{id}', [BeritaController::class, 'updateBerita'])->name('edit.berita');
Route::get('create_berita', [BeritaController::class, 'create']);
Route::post('addberita', [BeritaController::class, 'addBerita'])->name('store.berita');


// PENGUMUMAN
Route::get('/pengumuman', [PengumumanController::class, 'getPengumumanApi'])->name('pengumuman');
Route::delete('/pengumuman/{id}', [PengumumanController::class, 'deletePengumuman'])->name('hapus.pengumuman');
Route::get('/pengumuman/{id}', [PengumumanController::class, 'showUpdate'])->name('show.pengumuman');
Route::put('/editpengumuman/{id}', [PengumumanController::class, 'updatePengumuman'])->name('edit.pengumuman');
Route::get('create_pengumuman', [PengumumanController::class, 'create']);
Route::post('addpengumuman', [PengumumanController::class, 'addPengumuman'])->name('store.pengumuman');


// ACARA
Route::get('/acara', [AcaraController::class, 'getAcaraApi'])->name('acara');
Route::delete('/acara/{id}', [AcaraController::class, 'deleteAcara'])->name('hapus.acara');
Route::get('/acara/{id}', [AcaraController::class, 'showUpdate'])->name('show.acara');
Route::put('/editacara/{id}', [AcaraController::class, 'updateAcara'])->name('edit.acara');
Route::get('create_acara', [AcaraController::class, 'create']);
Route::post('addacara', [AcaraController::class, 'addAcara'])->name('store.acara');


// RENUNGAN
Route::get('/renungan', [RenunganController::class, 'getRenunganApi'])->name('renungan');
Route::delete('/renungan/{id}', [RenunganController::class, 'deleteRenungan'])->name('hapus.renungan');
Route::get('/renungan/{id}', [RenunganController::class, 'showUpdate'])->name('show.renungan');
Route::put('/editrenungan/{id}', [RenunganController::class, 'updateRenungan'])->name('edit.renungan');
Route::get('create_renungan', [RenunganController::class, 'create']);
Route::post('addrenungan', [RenunganController::class, 'addRenungan'])->name('store.renungan');


// SEKTOR
Route::get('/sektor', [SektorController::class, 'getSektorApi'])->name('sektor');
Route::delete('/sektor/{id}', [SektorController::class, 'deleteSektor'])->name('hapus.sektor');
Route::get('/sektor/{id}', [SektorController::class, 'showUpdate'])->name('show.sektor');
Route::put('/editsektor/{id}', [SektorController::class, 'updateSektor'])->name('edit.sektor');
Route::get('create_sektor', [SektorController::class, 'create']);
Route::post('addsektor', [SektorController::class, 'addSektor'])->name('store.sektor');


// PARHALADO
Route::get('/parhalado', [ParhaladoController::class, 'getParhaladoApi'])->name('parhalado');
Route::delete('/parhalado/{id}', [ParhaladoController::class, 'deleteParhalado'])->name('hapus.parhalado');
Route::get('/parhalado/{id}', [ParhaladoController::class, 'showUpdate'])->name('show.parhalado');
Route::put('/editparhalado/{id}', [ParhaladoController::class, 'updateParhalado'])->name('edit.parhalado');
Route::get('create_parhalado', [ParhaladoController::class, 'create']);
Route::post('addparhalado', [ParhaladoController::class, 'addParhalado'])->name('store.parhalado');
