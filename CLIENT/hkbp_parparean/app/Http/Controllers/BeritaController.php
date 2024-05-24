<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Http;

class BeritaController extends Controller
{

    public function getBeritaApi()
    {

        $client = new Client();
        $url = "http://localhost:9011/berita";
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('berita', [
            'data' => $contentArray,
            'title' => 'Berita'

        ]);
    }



    public function create()
    {
        return view('create_berita');
    }


    public function addBerita(Request $request)
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
        $response = Http::post("http://localhost:9011/berita", [
            'judul' => $request->input('judul'),
            'tanggal' => $request->input('tanggal'),
            'keterangan' => $request->input('keterangan'),
            'status' => $request->input('status'),
            // tambahkan data lainnya yang diperlukan untuk registrasi
        ]);


        // Memeriksa respons dari API
        if ($response->successful()) {
            // Registrasi berhasil, arahkan pengguna ke halaman login atau halaman setelah registrasi
            return redirect()->route('berita')->with('success', 'Data berhasil ditambah');
            // print_r($response);
        } else {
            // Registrasi gagal, kembalikan pengguna ke halaman registrasi dengan pesan kesalahan
            return response()->json([
                'message' => 'Gagal menambahkan data'
            ], 404);
        }
    }



    public function deleteBerita($id)
    {
        // Mengirimkan permintaan DELETE ke endpoint DeleteBerita pada API Go.
        $response = Http::delete('http://localhost:9011/berita/' . $id);

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
        $url = "http://localhost:9011/berita/" . $id;
        $response = $client->request('GET', $url);
        // // dd($response);
        // // echo $response->getBody()->getContents();

        $content = $response->getBody()->getContents();
        $contentArray = json_decode($content, true);
        // print_r($contentArray);
        return view('edit_berita', [
            'data' => $contentArray,
            'title' => 'berita'

        ]);
    }

    public function updateBerita(Request $request, $id)
    {


        // $request->validate([
        //     'judul' => ['sometimes', 'required', 'string', 'max:255'],
        //     'tanggal' => ['sometimes', 'required', 'string', 'max:255'],
        //     'keterangan' => ['sometimes', 'required', 'string', 'max:255'],
        //     'status' => ['sometimes', 'required', 'string', 'max:255'],
        // ]);

        $response = Http::put('http://localhost:9011/berita/' . $id, [
            'judul' => $request->input('judul'),
            'tanggal' => $request->input('tanggal'),
            'keterangan' => $request->input('keterangan'),
            'status' => $request->input('status'),
        ]);
        // Memeriksa apakah permintaan berhasil atau tidak
        if ($response->successful()) {
            return redirect()->route('berita')->with('success', 'Data berhasil diupdate');
        } else {
            return redirect()->back()->with('error', 'Gagal memperbarui data');
        }
    }
}
