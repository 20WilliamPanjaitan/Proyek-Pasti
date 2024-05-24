<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Http;

class SektorController extends Controller
{

    public function getSektorApi()
    {

        $client = new Client();
        $url = "http://localhost:9013/sektor";
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('sektor', [
            'data' => $contentArray,
            'title' => 'Sektor'

        ]);
    }



    public function create()
    {
        return view('create_sektor');
    }


    public function addSektor(Request $request)
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
        $response = Http::post("http://localhost:9013/sektor", [
            'nama_sektor' => $request->input('nama_sektor'),
            'deskripsi' => $request->input('deskripsi'),
            'alamat' => $request->input('alamat'),
            'kontak' => $request->input('kontak'),
            // tambahkan data lainnya yang diperlukan untuk registrasi
        ]);


        // Memeriksa respons dari API
        if ($response->successful()) {
            // Registrasi berhasil, arahkan pengguna ke halaman login atau halaman setelah registrasi
            return redirect()->route('sektor')->with('success', 'Data berhasil ditambah');
            // print_r($response);
        } else {
            // Registrasi gagal, kembalikan pengguna ke halaman registrasi dengan pesan kesalahan
            return response()->json([
                'message' => 'Gagal menambahkan data'
            ], 404);
        }
    }



    public function deleteSektor($id)
    {
        // Mengirimkan permintaan DELETE ke endpoint DeleteSektor pada API Go.
        $response = Http::delete('http://localhost:9013/sektor/' . $id);

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
        $url = "http://localhost:9013/sektor/" . $id;
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('edit_sektor', [
            'data' => $contentArray,
            'title' => 'sektor'

        ]);
    }

    public function updateSektor(Request $request, $id)
    {
        // $request->validate([
        //     'judul' => ['sometimes', 'required', 'string', 'max:255'],
        //     'tanggal' => ['sometimes', 'required', 'string', 'max:255'],
        //     'keterangan' => ['sometimes', 'required', 'string', 'max:255'],
        //     'status' => ['sometimes', 'required', 'string', 'max:255'],
        // ]);

        $response = Http::put('http://localhost:9013/sektor/' . $id, [
            'nama_sektor' => $request->input('nama_sektor'),
            'deskripsi' => $request->input('deskripsi'),
            'alamat' => $request->input('alamat'),
            'kontak' => $request->input('kontak'),
        ]);
        // Memeriksa apakah permintaan berhasil atau tidak
        if ($response->successful()) {
            return redirect()->route('sektor')->with('success', 'Data berhasil diupdate');
        } else {
            return redirect()->back()->with('error', 'Gagal memperbarui data');
        }
    }
}
