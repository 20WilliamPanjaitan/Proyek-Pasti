<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Http;

class AcaraController extends Controller
{

    public function getAcaraApi()
    {

        $client = new Client();
        $url = "http://localhost:9009/acara";
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('acara', [
            'data' => $contentArray,
            'title' => 'Acara'

        ]);
    }



    public function create()
    {
        return view('create_acara');
    }


    public function addAcara(Request $request)
    {
        // dd($request->status);
        // Validasi input pengguna
        // $request->validate([
        //     'judul' => ['sometimes', 'required', 'string', 'max:255'],
        //     'tanggal' => ['sometimes', 'required', 'string', 'max:255'],
        //     'keterangan' => ['sometimes', 'required', 'string', 'max:255'],
        //     'status' => ['sometimes', 'required', 'string', 'max:255'],
        //     // tambahkan validasi lainnya sesuai kebutuhan
        // ]);

        // Kirim data ke API eksternal untuk registrasi
        // dd($request);
        $response = Http::post("http://localhost:9009/acara", [
            'nama_acara' => $request->input('nama_acara'),
            'lokasi_acara' => $request->input('lokasi_acara'),
            'jenis_acara' => $request->input('jenis_acara'),
            'waktu_pelaksanaan' => $request->input('waktu_pelaksanaan'),
            'keterangan' => $request->input('keterangan'),
            // tambahkan data lainnya yang diperlukan untuk registrasi
        ]);


        // Memeriksa respons dari API
        if ($response->successful()) {
            // Registrasi berhasil, arahkan pengguna ke halaman login atau halaman setelah registrasi
            return redirect()->route('acara')->with('success', 'Data berhasil ditambah');
            // print_r($response);
        } else {
            // Registrasi gagal, kembalikan pengguna ke halaman registrasi dengan pesan kesalahan
            return response()->json([
                'message' => 'Gagal menambahkan data'
            ], 404);
        }
    }



    public function deleteAcara($id)
    {
        // Mengirimkan permintaan DELETE ke endpoint DeleteAcara pada API Go.
        $response = Http::delete('http://localhost:9009/acara/' . $id);

        // Mendapatkan respon dari API.
        $responseData = $response->json();

        // Melakukan penanganan sesuai dengan respon dari API.
        if ($response->successful()) {
            // Tindakan yang diambil jika permintaan berhasil (status kode 2xx).
            return redirect()->back();
        } else {
            // Tindakan yang diambil jika permintaan gagal (status kode selain 2xx).
            return redirect()->back();
        }
    }



    public function showUpdate($id)
    {
        $client = new Client();
        $url = "http://localhost:9009/acara/" . $id;
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('edit_acara', [
            'data' => $contentArray,
            'title' => 'acara'

        ]);
    }


    public function updateAcara(Request $request, $id)
    {


        // $request->validate([
        //     'judul' => ['sometimes', 'required', 'string', 'max:255'],
        //     'tanggal' => ['sometimes', 'required', 'string', 'max:255'],
        //     'keterangan' => ['sometimes', 'required', 'string', 'max:255'],
        //     'status' => ['sometimes', 'required', 'string', 'max:255'],
        // ]);

        $response = Http::put('http://localhost:9009/acara/' . $id, [
            'nama_acara' => $request->input('nama_acara'),
            'lokasi_acara' => $request->input('lokasi_acara'),
            'jenis_acara' => $request->input('jenis_acara'),
            'waktu_pelaksanaan' => $request->input('waktu_pelaksanaan'),
            'keterangan' => $request->input('keterangan'),
        ]);
        // Memeriksa apakah permintaan berhasil atau tidak
        if ($response->successful()) {
            return redirect()->route('acara')->with('success', 'Data berhasil diupdate');
        } else {
            return redirect()->back()->with('error', 'Gagal memperbarui data');
        }
    }
}
